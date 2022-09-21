package class

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
	ctx          = context.Background()
	classService = NewService(factory.NewFactory())
)

func TestClassServiceFindAllSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	var (
		asserts = assert.New(t)
		payload = pkgdto.SearchGetRequest{}
	)

	res, err := classService.Find(ctx, &payload)
	if err != nil {
		t.Fatal(err)
	}

	asserts.Len(res.Data, 2)
	for _, val := range res.Data {
		asserts.NotEmpty(val.Name)
		asserts.NotEmpty(val.ID)
	}
}
func TestClassServiceFindByIdSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	var (
		asserts = assert.New(t)
		payload = pkgdto.ByIDRequest{ID: 1}
	)

	res, err := classService.FindByID(ctx, &payload)
	if err != nil {
		t.Fatal(err)
	}

	asserts.Equal(uint(1), res.ID)
}

func TestClassServiceFindByIdRecordNotFound(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()

	var (
		asserts = assert.New(t)
		payload = pkgdto.ByIDRequest{ID: 1}
	)

	_, err := classService.FindByID(ctx, &payload)
	if err != nil {
		asserts.Equal(err.Error(), "error code 404")
	}
}

func TestClassServiceUpdateByIdSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	var (
		asserts = assert.New(t)
		id      = uint(1)
		name    = "Finance Dept."
		payload = dto.UpdateClassRequestBody{
			ID:   &id,
			Name: &name,
		}
	)
	res, err := classService.UpdateById(ctx, &payload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.Equal(name, res.Name)
}

func TestClassServiceUpdateByIdRecordNotFound(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()

	var (
		asserts = assert.New(t)
		id      = uint(1)
		name    = "Finance Dept."
		payload = dto.UpdateClassRequestBody{
			ID:   &id,
			Name: &name,
		}
	)

	_, err := classService.UpdateById(ctx, &payload)
	if err != nil {
		asserts.Equal(err.Error(), "error code 404")
	}
}

func TestClassServiceDeleteByIdSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	var (
		asserts = assert.New(t)
		id      = uint(1)
		payload = pkgdto.ByIDRequest{ID: id}
	)

	res, err := classService.DeleteById(ctx, &payload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.NotNil(res.DeletedAt)
}

func TestClassServiceDeleteByIdRecordNotFound(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()

	var (
		asserts = assert.New(t)
		id      = uint(10)
		payload = pkgdto.ByIDRequest{ID: id}
	)

	_, err := classService.DeleteById(ctx, &payload)
	if asserts.Error(err) {
		asserts.Equal(err.Error(), "error code 404")
	}
}

func TestClassServiceCreateClassSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()

	var (
		asserts = assert.New(t)
		name    = "Finance Dept."
		payload = dto.CreateClassRequestBody{
			Name: &name,
		}
	)

	res, err := classService.Store(ctx, &payload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.NotEmpty(res.ID)
	asserts.Equal(*payload.Name, res.Name)
}

func TestClassServiceCreateClassAlreadyExist(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	var (
		asserts = assert.New(t)
		name    = enum.Class(testAClassID).String()
		payload = dto.CreateClassRequestBody{
			Name: &name,
		}
	)

	_, err := classService.Store(ctx, &payload)
	if asserts.Error(err) {
		asserts.Equal(err.Error(), "error code 409")
	}
}
