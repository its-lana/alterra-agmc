package student

import (
	"context"
	"testing"

	"student-service/database"
	"student-service/database/seeder"
	"student-service/internal/dto"
	"student-service/internal/factory"
	"student-service/internal/pkg/enum"
	pkgdto "student-service/pkg/dto"

	"github.com/stretchr/testify/assert"
)

var (
	testID                   = uint(1)
	testFullname             = "Vincent Luis Hubbard"
	ctx                      = context.Background()
	testStudentService       = NewService(factory.NewFactory())
	testUpdateStudentPayload = dto.UpdateStudentRequestBody{
		ID:       &testID,
		Fullname: &testFullname,
		Email:    &testEmail,
		MajorID:  &testMajorID,
		ClassID:  &testAClassID,
	}
	testFindAllPayload  = pkgdto.SearchGetRequest{}
	testFindByIdPayload = pkgdto.ByIDRequest{ID: 1}
)

func TestStudentServiceFindAllSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	asserts := assert.New(t)
	res, err := testStudentService.Find(ctx, &testFindAllPayload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.Len(res.Data, 3)
	for _, val := range res.Data {
		asserts.NotEmpty(val.Email)
		asserts.NotEmpty(val.Fullname)
		asserts.NotEmpty(val.ID)
	}
}

func TestStudentServiceFindByIdSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	asserts := assert.New(t)
	res, err := testStudentService.FindByID(ctx, &testFindByIdPayload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.Equal(uint(1), res.ID)
}

func TestStudentServiceFindByIdRecordNotFound(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()

	asserts := assert.New(t)
	_, err := testStudentService.FindByID(ctx, &testFindByIdPayload)
	if err != nil {
		asserts.Equal(err.Error(), "error code 404")
	}
}

func TestStudentServiceUpdateByIdSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	asserts := assert.New(t)
	res, err := testStudentService.UpdateById(ctx, &testUpdateStudentPayload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.Equal(testFullname, res.Fullname)
	asserts.Equal(testEmail, res.Email)
	asserts.Equal(enum.Major(testMajorID).String(), res.Major.Name)
	asserts.Equal(enum.Class(testAClassID).String(), res.Class.Name)
}

func TestStudentServiceUpdateByIdRecordNotFound(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()

	asserts := assert.New(t)
	_, err := testStudentService.UpdateById(ctx, &testUpdateStudentPayload)
	if err != nil {
		asserts.Equal(err.Error(), "error code 404")
	}
}

func TestStudentServiceDeleteByIdSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	asserts := assert.New(t)
	res, err := testStudentService.DeleteById(ctx, &testFindByIdPayload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.NotNil(res.DeletedAt)
}

func TestStudentServiceDeleteByIdRecordNotFound(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()

	asserts := assert.New(t)
	_, err := testStudentService.DeleteById(ctx, &testFindByIdPayload)
	if asserts.Error(err) {
		asserts.Equal(err.Error(), "error code 404")
	}
}
