package controllers

import (
	"Task-Code/config"
	"Task-Code/lib/database/seeder"
	"Task-Code/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// init function testing
func setupBookTest(t *testing.T) {
	//load env
	if err := godotenv.Load("../.env"); err != nil {
		t.Error("Error loading .env file")
	}

	//setup database
	config.InitDB()

	// clear database
	s := seeder.NewSeeder()
	fmt.Println(s)
	s.BookDelete()
	s.BookSeed()
}

func TestGetAllBooksSuccess(t *testing.T) {
	setupBookTest(t)
	//setup echo context
	e := echo.New()

	//setup request
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//test
	assert.NoError(t, GetAllBooks(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	bodyRes, _ := io.ReadAll(rec.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bodyRes, &responseBody)

	assert.Equal(t, "Success Get Books", responseBody["status"])
}

func TestGetAllBooksFailedDBNotConnect(t *testing.T) {
	setupBookTest(t)
	db, err := config.DB.DB()
	assert.NoError(t, err)
	assert.NoError(t, db.Close())
	//setup echo context
	e := echo.New()

	//setup request
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//test
	err = GetAllBooks(c)
	assert.Error(t, err)
	hErr, ok := err.(*echo.HTTPError)
	assert.True(t, ok)
	assert.Equal(t, http.StatusInternalServerError, hErr.Code)
}

func TestAddNewBooksSuccess(t *testing.T) {
	setupBookTest(t)

	//setup echo context
	e := echo.New()

	//create json body
	body := models.Book{
		Title:     "Test Book 3",
		Author:    "Tester 3",
		TotalPage: 740,
	}

	//setup request
	b, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(string(b)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//test
	assert.NoError(t, AddNewBook(c))
	assert.Equal(t, http.StatusOK, rec.Code)

	bodyRes, _ := io.ReadAll(rec.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bodyRes, &responseBody)

	assert.Equal(t, "New Book Added Successfully!", responseBody["status"])
}

func TestAddNewBooksFailedWhenUserNotInputAuthor(t *testing.T) {
	setupBookTest(t)

	//setup echo context
	e := echo.New()

	//create json body
	body := models.Book{
		Title:     "Test Book 3",
		Author:    "",
		TotalPage: 740,
	}
	b, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(string(b)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//test
	assert.NoError(t, AddNewBook(c))
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	bodyRes, _ := io.ReadAll(rec.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bodyRes, &responseBody)

	assert.Equal(t, "author is required", responseBody["status"])

}

func TestGetBookByIdSuccess(t *testing.T) {
	setupBookTest(t)

	//setup echo context
	e := echo.New()

	//setup request
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//set params
	c.SetParamNames("id")
	c.SetParamValues("1")

	//test
	assert.NoError(t, GetBookByID(c))
	assert.Equal(t, http.StatusOK, rec.Code)

	bodyRes, _ := io.ReadAll(rec.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bodyRes, &responseBody)

	assert.Equal(t, "Success Get Book By Id", responseBody["status"])
}

func TestGetBookByIdNotFound(t *testing.T) {
	setupBookTest(t)

	//setup echo context
	e := echo.New()

	//setup request
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//set params
	c.SetParamNames("id")
	c.SetParamValues("10")

	//test
	assert.NoError(t, GetBookByID(c))
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	bodyRes, _ := io.ReadAll(rec.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bodyRes, &responseBody)

	assert.Equal(t, "Book not found!", responseBody["status"])
}

func TestUpdateBookByIdSuccess(t *testing.T) {
	setupBookTest(t)

	//setup echo context
	e := echo.New()

	//create json body
	body := models.Book{
		Title:     "Test Book Z",
		Author:    "Tester Z",
		TotalPage: 940,
	}

	//setup request
	b, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPut, "/books", strings.NewReader(string(b)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//set params
	c.SetParamNames("id")
	c.SetParamValues("1")

	//test
	assert.NoError(t, UpdateBook(c))
	assert.Equal(t, http.StatusOK, rec.Code)

	bodyRes, _ := io.ReadAll(rec.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bodyRes, &responseBody)

	assert.Equal(t, "Book Updated Successfully!", responseBody["status"])
}

func TestUpdateBookByIdNotFound(t *testing.T) {
	setupBookTest(t)

	//setup echo context
	e := echo.New()

	//create json body
	body := models.Book{
		Title:     "Test Book Z",
		Author:    "Tester Z",
		TotalPage: 940,
	}

	//setup request
	b, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPut, "/books", strings.NewReader(string(b)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//set params
	c.SetParamNames("id")
	c.SetParamValues("100")

	//test
	assert.NoError(t, UpdateBook(c))
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	bodyRes, _ := io.ReadAll(rec.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bodyRes, &responseBody)

	assert.Equal(t, fmt.Sprintf("Update Failed! Book Id not found!"), responseBody["status"])
}

func TestDeleteBookByIdSuccess(t *testing.T) {
	setupBookTest(t)

	//setup echo context
	e := echo.New()

	//setup request
	req := httptest.NewRequest(http.MethodDelete, "/books", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//set params
	c.SetParamNames("id")
	c.SetParamValues("1")

	//test
	assert.NoError(t, DeleteBook(c))
	assert.Equal(t, http.StatusOK, rec.Code)

	bodyRes, _ := io.ReadAll(rec.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bodyRes, &responseBody)

	assert.Equal(t, "Book Deleted Successfully!", responseBody["status"])
}

func TestDeleteBookByIdNotFound(t *testing.T) {
	setupBookTest(t)

	//setup echo context
	e := echo.New()

	//setup request
	req := httptest.NewRequest(http.MethodDelete, "/books", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//set params
	c.SetParamNames("id")
	c.SetParamValues("10")

	//set user id
	c.Set("userId", 10)

	//test
	assert.NoError(t, DeleteBook(c))
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	bodyRes, _ := io.ReadAll(rec.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bodyRes, &responseBody)

	assert.Equal(t, fmt.Sprintf("Delete Failed! Book Id not found!"), responseBody["status"])
}
