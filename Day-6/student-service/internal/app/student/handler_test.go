package student

import (
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"student-service/database"
	"student-service/database/seeder"
	"student-service/internal/factory"
	"student-service/internal/mocks"
	"student-service/internal/pkg/enum"
	"student-service/internal/pkg/util"
	"student-service/internal/repository"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	adminClaims    = util.CreateJWTClaims(testEmail, testStudentID, testAClassID, testMajorID)
	userClaims     = util.CreateJWTClaims(testEmail, uint(2), uint(enum.B), testMajorID)
	db             = database.GetConnection()
	echoMock       = mocks.EchoMock{E: echo.New()}
	studentHandler = NewHandler(&f)
	f              = factory.Factory{StudentRepository: repository.NewStudentRepository(db)}
	testAClassID   = uint(enum.A)
	testMajorID    = uint(enum.Finance)
	testEmail      = "vincentlhubbard@edu.ac.id"
	testStudentID  = uint(1)
)

func TestStudentHandlerGetInvalidPayload(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	token, err := util.CreateJWTToken(adminClaims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/api/v1/students")
	c.QueryParams().Add("page", "a")
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(studentHandler.Get(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Bad Request")
	}
}

func TestStudentHandlerGetUnauthorized(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	c.SetPath("/api/v1/students")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(studentHandler.Get(c)) {
		asserts.Equal(401, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "unauthorized")
	}
}

func TestStudentHandlerGetSuccess(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	token, err := util.CreateJWTToken(adminClaims)
	if err != nil {
		t.Fatal(err)
	}
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	c.SetPath("/api/v1/students")
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(studentHandler.Get(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "id")
		asserts.Contains(body, "fullname")
		asserts.Contains(body, "email")
	}
}

func TestStudentHandlerGetByIdInvalidPayload(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	studentID := "a"

	token, err := util.CreateJWTToken(adminClaims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/api/v1/students")
	c.SetParamNames("id")
	c.SetParamValues(studentID)
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(studentHandler.GetById(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Bad Request")
	}
}

func TestStudentHandlerGetByIdNotFound(t *testing.T) {
	seeder.NewSeeder().DeleteAll()

	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	token, err := util.CreateJWTToken(adminClaims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/api/v1/students")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(testStudentID)))
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(studentHandler.GetById(c)) {
		asserts.Equal(404, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Data not found")
	}
}

func TestStudentHandlerGetByIdUnauthorized(t *testing.T) {
	seeder.NewSeeder().DeleteAll()

	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)

	c.SetPath("/api/v1/students")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(testStudentID)))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(studentHandler.GetById(c)) {
		asserts.Equal(401, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "unauthorized")
	}
}

func TestStudentHandlerGetByIdSuccess(t *testing.T) {
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	token, err := util.CreateJWTToken(adminClaims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/api/v1/students")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(testStudentID)))
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(studentHandler.GetById(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "id")
		asserts.Contains(body, "fullname")
		asserts.Contains(body, "email")
		asserts.Contains(body, "class")
		asserts.Contains(body, "major")
	}
}

func TestStudentHandlerUpdateByIdInvalidPayload(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodPut, "/", nil)
	studentID := "a"
	token, err := util.CreateJWTToken(adminClaims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/api/v1/students")
	c.SetParamNames("id")
	c.SetParamValues(studentID)
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(studentHandler.UpdateById(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Bad Request")
	}
}

func TestStudentHandlerUpdateByIdNotFound(t *testing.T) {
	seeder.NewSeeder().DeleteAll()

	c, rec := echoMock.RequestMock(http.MethodPut, "/", nil)
	token, err := util.CreateJWTToken(adminClaims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/api/v1/students")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(testStudentID)))
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(studentHandler.UpdateById(c)) {
		asserts.Equal(404, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Data not found")
	}
}

func TestStudentHandlerUpdateByIdUnauthorized(t *testing.T) {
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	c, rec := echoMock.RequestMock(http.MethodPut, "/", nil)
	token, err := util.CreateJWTToken(userClaims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/api/v1/students")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(3)))
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(studentHandler.UpdateById(c)) {
		asserts.Equal(401, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "unauthorized")
	}
}

func TestStudentHandlerUpdateByIdSuccess(t *testing.T) {
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	c, rec := echoMock.RequestMock(http.MethodPut, "/", nil)
	token, err := util.CreateJWTToken(userClaims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/api/v1/students")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(userClaims.BID)))
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(studentHandler.UpdateById(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "id")
		asserts.Contains(body, "fullname")
		asserts.Contains(body, "email")
		asserts.Contains(body, "class")
		asserts.Contains(body, "major")
	}
}

func TestStudentHandlerDeleteByIdInvalidPayload(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodDelete, "/", nil)
	studentID := "a"
	token, err := util.CreateJWTToken(adminClaims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/api/v1/students")
	c.SetParamNames("id")
	c.SetParamValues(studentID)
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(studentHandler.DeleteById(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Bad Request")
	}
}

func TestStudentHandlerDeleteByIdNotFound(t *testing.T) {
	seeder.NewSeeder().DeleteAll()

	c, rec := echoMock.RequestMock(http.MethodDelete, "/", nil)
	token, err := util.CreateJWTToken(adminClaims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/api/v1/students")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(testStudentID)))
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(studentHandler.DeleteById(c)) {
		asserts.Equal(404, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Data not found")
	}
}

func TestStudentHandlerDeleteByIdUnauthorized(t *testing.T) {
	seeder.NewSeeder().DeleteAll()

	c, rec := echoMock.RequestMock(http.MethodDelete, "/", nil)

	c.SetPath("/api/v1/students")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(testStudentID)))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(studentHandler.DeleteById(c)) {
		asserts.Equal(401, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "unauthorized")
	}
}

func TestStudentHandlerDeleteByIdSuccess(t *testing.T) {
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	c, rec := echoMock.RequestMock(http.MethodDelete, "/", nil)
	token, err := util.CreateJWTToken(adminClaims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/api/v1/students")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(testStudentID)))
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(studentHandler.DeleteById(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "id")
		asserts.Contains(body, "fullname")
		asserts.Contains(body, "email")
		asserts.Contains(body, "deleted_at")
	}
}
