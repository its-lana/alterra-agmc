package major

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"student-service/database"
	"student-service/database/seeder"
	"student-service/internal/dto"
	"student-service/internal/factory"
	"student-service/internal/mocks"
	"student-service/internal/pkg/enum"
	"student-service/internal/pkg/util"
	"student-service/internal/repository"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	adminClaims       = util.CreateJWTClaims(testEmail, testStudentID, testAClassID, testMajorID)
	db                = database.GetConnection()
	majorHandler      = NewHandler(&f)
	echoMock          = mocks.EchoMock{E: echo.New()}
	f                 = factory.Factory{MajorRepository: repository.NewMajorRepository(db)}
	testAClassID      = uint(enum.A)
	testCreatePayload = dto.CreateMajorRequestBody{Name: &testMajorName}
	testMajorID       = uint(enum.Finance)
	testMajorName     = enum.Major(testMajorID).String()
	testEmail         = "vincentlhubbard@edu.ac.id"
	testStudentID     = uint(1)
	testUpdatePayload = dto.UpdateMajorRequestBody{ID: &testMajorID, Name: &testMajorName}
	testBClassID      = uint(enum.B)
	userClaims        = util.CreateJWTClaims(testEmail, testStudentID, testBClassID, testMajorID)
)

func TestMajorHandlerGetInvalidPayload(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	token, err := util.CreateJWTToken(adminClaims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/api/v1/majors")
	c.QueryParams().Add("page", "a")
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(majorHandler.Get(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Bad Request")
	}
}
func TestMajorHandlerGetUnauthorized(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	c.SetPath("/api/v1/majors")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(majorHandler.Get(c)) {
		asserts.Equal(401, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "unauthorized")
	}
}

func TestMajorHandlerGetSuccess(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	token, err := util.CreateJWTToken(adminClaims)
	if err != nil {
		t.Fatal(err)
	}
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	c.SetPath("/api/v1/majors")
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(majorHandler.Get(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "id")
		asserts.Contains(body, "name")
	}
}

func TestMajorHandlerGetByIdInvalidPayload(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	majorID := "a"

	token, err := util.CreateJWTToken(adminClaims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/api/v1/majors")
	c.SetParamNames("id")
	c.SetParamValues(majorID)
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(majorHandler.GetById(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Bad Request")
	}
}

func TestMajorHandlerGetByIdNotFound(t *testing.T) {
	seeder.NewSeeder().DeleteAll()

	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	majorID := strconv.Itoa(int(testMajorID))
	token, err := util.CreateJWTToken(adminClaims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/api/v1/majors")
	c.SetParamNames("id")
	c.SetParamValues(majorID)
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(majorHandler.GetById(c)) {
		asserts.Equal(404, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Data not found")
	}
}

func TestMajorHandlerGetByIdUnauthorized(t *testing.T) {
	seeder.NewSeeder().DeleteAll()

	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	majorID := strconv.Itoa(int(testMajorID))

	c.SetPath("/api/v1/majors")
	c.SetParamNames("id")
	c.SetParamValues(majorID)

	// testing
	asserts := assert.New(t)
	if asserts.NoError(majorHandler.GetById(c)) {
		asserts.Equal(401, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "unauthorized")
	}
}

func TestMajorHandlerGetByIdSuccess(t *testing.T) {
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	majorID := strconv.Itoa(int(testMajorID))
	token, err := util.CreateJWTToken(adminClaims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/api/v1/majors")
	c.SetParamNames("id")
	c.SetParamValues(majorID)
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(majorHandler.GetById(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "id")
		asserts.Contains(body, "name")
	}
}

func TestMajorHandlerUpdateByIdInvalidPayload(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodPut, "/", nil)
	majorID := "a"
	token, err := util.CreateJWTToken(adminClaims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/api/v1/majors")
	c.SetParamNames("id")
	c.SetParamValues(majorID)
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(majorHandler.UpdateById(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Bad Request")
	}
}

func TestMajorHandlerUpdateByIdNotFound(t *testing.T) {
	seeder.NewSeeder().DeleteAll()

	payload, err := json.Marshal(testUpdatePayload)
	if err != nil {
		t.Fatal(err)
	}
	c, rec := echoMock.RequestMock(http.MethodPut, "/", bytes.NewBuffer(payload))
	majorID := strconv.Itoa(int(testMajorID))
	token, err := util.CreateJWTToken(adminClaims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/api/v1/majors")
	c.SetParamNames("id")
	c.SetParamValues(majorID)
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	c.Request().Header.Set("Content-Type", "application/json")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(majorHandler.UpdateById(c)) {
		asserts.Equal(404, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Data not found")
	}
}
func TestMajorHandlerUpdateByIdUnauthorized(t *testing.T) {
	seeder.NewSeeder().DeleteAll()

	c, rec := echoMock.RequestMock(http.MethodPut, "/", nil)
	majorID := strconv.Itoa(int(testMajorID))

	c.SetPath("/api/v1/majors")
	c.SetParamNames("id")
	c.SetParamValues(majorID)

	token, err := util.CreateJWTToken(userClaims)
	if err != nil {
		t.Fatal(err)
	}
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(majorHandler.UpdateById(c)) {
		asserts.Equal(401, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "unauthorized")
	}
}

func TestMajorHandlerUpdateByIdSuccess(t *testing.T) {
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	payload, err := json.Marshal(testUpdatePayload)
	if err != nil {
		t.Fatal(err)
	}
	c, rec := echoMock.RequestMock(http.MethodPut, "/", bytes.NewBuffer(payload))
	majorID := strconv.Itoa(int(testMajorID))
	token, err := util.CreateJWTToken(adminClaims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/api/v1/majors")
	c.SetParamNames("id")
	c.SetParamValues(majorID)
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	c.Request().Header.Set("Content-Type", "application/json")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(majorHandler.UpdateById(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "id")
		asserts.Contains(body, "name")
	}
}

func TestMajorHandlerDeleteByIdInvalidPayload(t *testing.T) {
	c, rec := echoMock.RequestMock(http.MethodDelete, "/", nil)
	majorID := "a"
	token, err := util.CreateJWTToken(adminClaims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/api/v1/majors")
	c.SetParamNames("id")
	c.SetParamValues(majorID)
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(majorHandler.DeleteById(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Bad Request")
	}
}

func TestMajorHandlerDeleteByIdNotFound(t *testing.T) {
	seeder.NewSeeder().DeleteAll()

	c, rec := echoMock.RequestMock(http.MethodDelete, "/", nil)
	majorID := strconv.Itoa(int(testMajorID))
	token, err := util.CreateJWTToken(adminClaims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/api/v1/majors")
	c.SetParamNames("id")
	c.SetParamValues(majorID)
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(majorHandler.DeleteById(c)) {
		asserts.Equal(404, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Data not found")
	}
}

func TestMajorHandlerDeleteByIdUnauthorized(t *testing.T) {
	seeder.NewSeeder().DeleteAll()

	c, rec := echoMock.RequestMock(http.MethodDelete, "/", nil)
	majorID := strconv.Itoa(int(testMajorID))

	c.SetPath("/api/v1/majors")
	c.SetParamNames("id")
	c.SetParamValues(majorID)

	token, err := util.CreateJWTToken(userClaims)
	if err != nil {
		t.Fatal(err)
	}
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(majorHandler.DeleteById(c)) {
		asserts.Equal(401, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "unauthorized")
	}
}

func TestMajorHandlerDeleteByIdSuccess(t *testing.T) {
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	c, rec := echoMock.RequestMock(http.MethodDelete, "/", nil)
	majorID := strconv.Itoa(int(testMajorID))
	token, err := util.CreateJWTToken(adminClaims)
	if err != nil {
		t.Fatal(err)
	}

	c.SetPath("/api/v1/majors")
	c.SetParamNames("id")
	c.SetParamValues(majorID)
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(majorHandler.DeleteById(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "id")
		asserts.Contains(body, "name")
		asserts.Contains(body, "deleted_at")
	}
}

func TestMajorHandlerCreateInvalidPayload(t *testing.T) {
	token, err := util.CreateJWTToken(adminClaims)
	if err != nil {
		t.Fatal(err)
	}
	payload, err := json.Marshal(dto.CreateMajorRequestBody{})
	if err != nil {
		t.Fatal(err)
	}

	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))

	c.SetPath("/api/v1/majors")
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	c.Request().Header.Set("Content-Type", "application/json")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(majorHandler.Create(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Invalid parameters or payload")
	}
}

func TestMajorHandlerCreateMajorAlreadyExist(t *testing.T) {
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	token, err := util.CreateJWTToken(adminClaims)
	if err != nil {
		t.Fatal(err)
	}
	name := enum.Major(testMajorID).String()
	payload, err := json.Marshal(dto.CreateMajorRequestBody{Name: &name})
	if err != nil {
		t.Fatal(err)
	}

	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))

	c.SetPath("/api/v1/majors")
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	c.Request().Header.Set("Content-Type", "application/json")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(majorHandler.Create(c)) {
		asserts.Equal(409, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "duplicate")
	}
}

func TestMajorHandlerCreateUnauthorized(t *testing.T) {
	payload, err := json.Marshal(testCreatePayload)
	if err != nil {
		t.Fatal(err)
	}

	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))

	c.SetPath("/api/v1/majors")
	c.Request().Header.Set("Content-Type", "application/json")

	token, err := util.CreateJWTToken(userClaims)
	if err != nil {
		t.Fatal(err)
	}
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// testing
	asserts := assert.New(t)
	if asserts.NoError(majorHandler.Create(c)) {
		asserts.Equal(401, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "unauthorized")
	}
}

func TestMajorHandlerCreateSuccess(t *testing.T) {
	seeder.NewSeeder().DeleteAll()

	token, err := util.CreateJWTToken(adminClaims)
	if err != nil {
		t.Fatal(err)
	}
	payload, err := json.Marshal(testCreatePayload)
	if err != nil {
		t.Fatal(err)
	}

	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))

	c.SetPath("/api/v1/majors")
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	c.Request().Header.Set("Content-Type", "application/json")

	// testing
	asserts := assert.New(t)
	if asserts.NoError(majorHandler.Create(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "id")
		asserts.Contains(body, "name")
	}
}
