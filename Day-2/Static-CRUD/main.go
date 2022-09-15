package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Book struct {
	Id        uint   `json:"id" form:"id"`
	Title     string `json:"title" form:"title"`
	Author    string `json:"author" form:"author"`
	TotalPage uint   `json:"total_page" form:"total_page"`
}

type HTTPResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

var book_collection = []Book{
	{Id: 1, Title: "Rindu", Author: "Tere Liye", TotalPage: 523},
	{Id: 2, Title: "Hujan", Author: "Tere Liye", TotalPage: 320},
	{Id: 3, Title: "Orang-orang Biasa", Author: "Andrea Hirata", TotalPage: 262},
}

func (resp *HTTPResponse) WriteToResponseBody(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Code)
	return json.NewEncoder(w).Encode(resp)
}

func WrapResponse(code int, status string, response interface{}) *HTTPResponse {
	newResponse := new(HTTPResponse)
	newResponse.Code = code
	newResponse.Status = status
	newResponse.Data = response
	return newResponse
}

func AddNewBook(c echo.Context) error {
	book := Book{}
	err := c.Bind(&book)
	if err != nil {
		return WrapResponse(http.StatusBadRequest, "Bad Request, Failed Add New Book!", &Book{}).WriteToResponseBody(c.Response())
	}
	book_collection = append(book_collection, book)

	return WrapResponse(http.StatusCreated, "New Book Added Succesfully!", &book).WriteToResponseBody(c.Response())
}

func GetBookById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, book := range book_collection {
		if book.Id == uint(id) {
			return WrapResponse(http.StatusOK, "Get Book by Id Succesfully!", book).WriteToResponseBody(c.Response())
		}
	}

	return WrapResponse(http.StatusBadRequest, "Failed Get Book by Id", &Book{}).WriteToResponseBody(c.Response())
}

func GetAllBooks(c echo.Context) error {
	return WrapResponse(http.StatusOK, "Get All Books Successfully!", book_collection).WriteToResponseBody(c.Response())
}

func UpdateBookById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	bookPayload := Book{}

	err := c.Bind(&bookPayload)

	if err != nil {
		return WrapResponse(http.StatusBadRequest, "Bad Request, Failed Update Book!", &bookPayload).WriteToResponseBody(c.Response())
	}
	for i, book := range book_collection {
		if book.Id == uint(id) {
			book_collection[i] = bookPayload
			return WrapResponse(http.StatusOK, "Book Updated Successfully!", &bookPayload).WriteToResponseBody(c.Response())
		}
	}

	return WrapResponse(http.StatusBadRequest, "Bad Request, Failed Update Book!", &bookPayload).WriteToResponseBody(c.Response())
}

func DeleteBookById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, book := range book_collection {
		if book.Id == uint(id) {
			book_collection[i] = book_collection[0]
			book_collection = book_collection[1:]
			return WrapResponse(http.StatusOK, "Book Deleted Successfully", &Book{}).WriteToResponseBody(c.Response())
		}
	}

	return WrapResponse(http.StatusBadRequest, "Failed Delete Book", &Book{}).WriteToResponseBody(c.Response())
}

func main() {
	e := echo.New()

	v1 := e.Group("/v1")
	books := v1.Group("/books")

	books.GET("", GetAllBooks)
	books.GET("/:id", GetBookById)
	books.POST("", AddNewBook)
	books.PUT("/:id", UpdateBookById)
	books.DELETE("/:id", DeleteBookById)

	e.Logger.Fatal(e.Start(":3000"))
}
