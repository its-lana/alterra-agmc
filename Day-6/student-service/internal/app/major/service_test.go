package major

import (
	"context"
	"testing"

	"student-service/database"
	"student-service/database/seeder"
	"student-service/internal/factory"
	pkgdto "student-service/pkg/dto"

	"github.com/stretchr/testify/assert"
)

var (
	ctx                 = context.Background()
	majorService        = NewService(factory.NewFactory())
	testFindAllPayload  = pkgdto.SearchGetRequest{}
	testFindByIdPayload = pkgdto.ByIDRequest{ID: 1}
)

func TestMajorServiceFindAllSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	asserts := assert.New(t)
	res, err := majorService.Find(ctx, &testFindAllPayload)
	if err != nil {
		t.Fatal(err)
	}

	asserts.Len(res.Data, 3)
	for _, val := range res.Data {
		asserts.NotEmpty(val.Name)
		asserts.NotEmpty(val.ID)
	}
}
func TestMajorServiceFindByIdSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	asserts := assert.New(t)
	res, err := majorService.FindByID(ctx, &testFindByIdPayload)
	if err != nil {
		t.Fatal(err)
	}

	asserts.Equal(uint(1), res.ID)
}

func TestMajorServiceFindByIdRecordNotFound(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()

	asserts := assert.New(t)
	_, err := majorService.FindByID(ctx, &testFindByIdPayload)
	if err != nil {
		asserts.Equal(err.Error(), "error code 404")
	}
}

func TestMajorServiceUpdateByIdSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	asserts := assert.New(t)
	res, err := majorService.UpdateById(ctx, &testUpdatePayload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.Equal(testMajorName, res.Name)
}

func TestMajorServiceUpdateByIdRecordNotFound(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()

	asserts := assert.New(t)
	_, err := majorService.UpdateById(ctx, &testUpdatePayload)
	if err != nil {
		asserts.Equal(err.Error(), "error code 404")
	}
}

func TestMajorServiceDeleteByIdSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	asserts := assert.New(t)
	res, err := majorService.DeleteById(ctx, &testFindByIdPayload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.NotNil(res.DeletedAt)
}

func TestMajorServiceDeleteByIdRecordNotFound(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()

	asserts := assert.New(t)
	_, err := majorService.DeleteById(ctx, &testFindByIdPayload)
	if asserts.Error(err) {
		asserts.Equal(err.Error(), "error code 404")
	}
}

func TestMajorServiceCreateMajorSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()

	asserts := assert.New(t)
	res, err := majorService.Store(ctx, &testCreatePayload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.NotEmpty(res.ID)
	asserts.Equal(*testCreatePayload.Name, res.Name)
}

func TestMajorServiceCreateMajorAlreadyExist(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	asserts := assert.New(t)
	_, err := majorService.Store(ctx, &testCreatePayload)
	if asserts.Error(err) {
		asserts.Equal(err.Error(), "error code 409")
	}
}
