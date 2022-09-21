package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"student-service/database"
	"student-service/database/seeder"
	"student-service/internal/dto"
	"student-service/internal/factory"
	"student-service/internal/mocks"
	"student-service/internal/repository"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestAuthHandlerLoginByEmailAndPasswordInvalidPayload(t *testing.T) {
	// setup database
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	// setup context
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	c, rec := echoMock.RequestMock(http.MethodPost, "/", nil)
	c.SetPath("/api/v1/auth/login")

	// setup handler
	asserts := assert.New(t)
	db := database.GetConnection()
	factory := factory.Factory{StudentRepository: repository.NewStudentRepository(db)}
	authHandler := NewHandler(&factory)

	// testing
	if asserts.NoError(authHandler.LoginByEmailAndPassword(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.JSONEq(`{"meta": {"success": false,"message": "Invalid parameters or payload","info": null},"error": "bad_request"}`, body)
	}
}

func TestAuthHandlerLoginByEmailAndPasswordUnmatchedEmailAndPassword(t *testing.T) {
	// setup database
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	// setup context
	emailAndPassword := dto.ByEmailAndPasswordRequest{
		Email:    "vincentlhubbard@edu.ac.id",
		Password: "1234567890",
	}
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	payload, err := json.Marshal(emailAndPassword)
	if err != nil {
		t.Fatal(err)
	}
	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))
	c.Request().Header.Set("Content-Type", "application/json")
	c.SetPath("/api/v1/auth/login")

	// setup handler
	asserts := assert.New(t)
	db := database.GetConnection()
	factory := factory.Factory{StudentRepository: repository.NewStudentRepository(db)}
	authHandler := NewHandler(&factory)

	// testing
	if asserts.NoError(authHandler.LoginByEmailAndPassword(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.JSONEq(`{"meta": {"success": false,"message": "Email or password is incorrect","info": null},"error": "bad_request"}`, body)
	}
}

func TestAuthHandlerLoginByEmailAndPasswordSuccess(t *testing.T) {
	// setup database
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	// setup context
	emailAndPassword := dto.ByEmailAndPasswordRequest{
		Email:    "vincentlhubbard@edu.ac.id",
		Password: "123abcABC!",
	}
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	payload, err := json.Marshal(emailAndPassword)
	if err != nil {
		t.Fatal(err)
	}
	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))
	c.Request().Header.Set("Content-Type", "application/json")
	c.SetPath("/api/v1/auth/login")

	// setup handler
	asserts := assert.New(t)
	db := database.GetConnection()
	factory := factory.Factory{StudentRepository: repository.NewStudentRepository(db)}
	authHandler := NewHandler(&factory)

	// testing
	if asserts.NoError(authHandler.LoginByEmailAndPassword(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "id")
		asserts.Contains(body, "fullname")
		asserts.Contains(body, "email")
		asserts.Contains(body, "jwt")
	}
}

func TestAuthHandlerRegisterByEmailAndPasswordBAlreadyExist(t *testing.T) {
	// setup database
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	// setup context
	var (
		fullname = "Vincent L. Hubbard"
		email    = "vincentlhubbard@edu.ac.id"
		password = "123abcABC!"
		majorID  = uint(1)
	)
	emailAndPassword := dto.RegisterStudentRequestBody{
		Fullname: fullname,
		Email:    email,
		Password: password,
		MajorID:  &majorID,
	}
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	payload, err := json.Marshal(emailAndPassword)
	if err != nil {
		t.Fatal(err)
	}
	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))
	c.Request().Header.Set("Content-Type", "application/json")
	c.SetPath("/api/v1/auth/signup")

	// setup handler
	asserts := assert.New(t)
	db := database.GetConnection()
	factory := factory.Factory{StudentRepository: repository.NewStudentRepository(db)}
	authHandler := NewHandler(&factory)

	// testing
	if asserts.NoError(authHandler.RegisterByEmailAndPassword(c)) {
		asserts.Equal(409, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Created value already exists")
	}
}

func TestAuthHandlerRegisterByEmailAndPasswordInvalidPayload(t *testing.T) {
	// setup database
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	// setup context
	var (
		fullname = "Vincent L. Hubbard"
		password = "123abcABC!"
		majorID  = uint(3)
	)
	emailAndPassword := dto.RegisterStudentRequestBody{
		Fullname: fullname,
		Password: password,
		MajorID:  &majorID,
	}
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	payload, err := json.Marshal(emailAndPassword)
	if err != nil {
		t.Fatal(err)
	}
	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))
	c.Request().Header.Set("Content-Type", "application/json")
	c.SetPath("/api/v1/auth/signup")

	// setup handler
	asserts := assert.New(t)
	db := database.GetConnection()
	factory := factory.Factory{StudentRepository: repository.NewStudentRepository(db)}
	authHandler := NewHandler(&factory)

	// testing
	if asserts.NoError(authHandler.RegisterByEmailAndPassword(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Invalid parameters or payload")
	}
}

func TestAuthHandlerRegisterByEmailAndPasswordSuccess(t *testing.T) {
	// setup database
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	// setup context
	var (
		fullname = "Azka"
		email    = "azka@edu.ac.id"
		password = "123abcABC!"
		majorID  = uint(2)
	)
	emailAndPassword := dto.RegisterStudentRequestBody{
		Fullname: fullname,
		Email:    email,
		Password: password,
		MajorID:  &majorID,
	}
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	payload, err := json.Marshal(emailAndPassword)
	if err != nil {
		t.Fatal(err)
	}
	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))
	c.Request().Header.Set("Content-Type", "application/json")
	c.SetPath("/api/v1/auth/signup")

	// setup handler
	asserts := assert.New(t)
	db := database.GetConnection()
	factory := factory.Factory{StudentRepository: repository.NewStudentRepository(db)}
	authHandler := NewHandler(&factory)

	// testing
	if asserts.NoError(authHandler.RegisterByEmailAndPassword(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "id")
		asserts.Contains(body, "fullname")
		asserts.Contains(body, "email")
		asserts.Contains(body, "jwt")
	}
}
