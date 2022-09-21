package auth

import (
	"context"
	"strings"
	"testing"

	"student-service/database"
	"student-service/database/seeder"
	"student-service/internal/dto"
	"student-service/internal/factory"

	"github.com/stretchr/testify/assert"
)

func TestAuthServiceLoginByEmailAndPasswordSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()
	asserts := assert.New(t)
	var (
		authService = NewService(factory.NewFactory())
		ctx         = context.Background()
		payload     = dto.ByEmailAndPasswordRequest{
			Email:    "vincentlhubbard@edu.ac.id",
			Password: "123abcABC!",
		}
	)
	res, err := authService.LoginByEmailAndPassword(ctx, &payload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.Equal(payload.Email, res.Email)
	asserts.Len(strings.Split(res.JWT, "."), 3)
}

func TestAuthServiceLoginByEmailAndPasswordRecordNotFound(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()
	var (
		asserts     = assert.New(t)
		authService = NewService(factory.NewFactory())
		ctx         = context.Background()
		payload     = dto.ByEmailAndPasswordRequest{
			Email:    "azkaframadhan@edu.ac.id",
			Password: "123abcABC!",
		}
	)
	_, err := authService.LoginByEmailAndPassword(ctx, &payload)
	if asserts.Error(err) {
		asserts.Equal(err.Error(), "error code 404")
	}
}

func TestAuthServiceLoginByEmailAndPasswordunmatchedEmailAndPassword(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()
	asserts := assert.New(t)
	var (
		authService = NewService(factory.NewFactory())
		ctx         = context.Background()
		payload     = dto.ByEmailAndPasswordRequest{
			Email:    "vincentlhubbard@edu.ac.id",
			Password: "1234567890",
		}
	)
	_, err := authService.LoginByEmailAndPassword(ctx, &payload)
	if asserts.Error(err) {
		asserts.Equal(err.Error(), "error code 400")
	}
}

func TestAuthServiceRegisterByEmailAndPasswordSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()
	asserts := assert.New(t)
	var (
		authService = NewService(factory.NewFactory())
		ctx         = context.Background()
		majorID     = uint(1)
		payload     = dto.RegisterStudentRequestBody{
			Fullname: "Azka Fadhli Ramadhan",
			Email:    "azkaframadhan@edu.ac.id",
			Password: "123abcABC!",
			MajorID:  &majorID,
		}
	)
	payload.FillDefaults()
	res, err := authService.RegisterByEmailAndPassword(ctx, &payload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.NotEmpty(res.ID)
	asserts.Equal(payload.Fullname, res.Fullname)
	asserts.Equal(payload.Email, res.Email)
	asserts.Len(strings.Split(res.JWT, "."), 3)
}

func TestAuthServiceRegisterByEmailAndPasswordBExist(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()
	asserts := assert.New(t)
	var (
		authService = NewService(factory.NewFactory())
		ctx         = context.Background()
		majorID     = uint(1)
		payload     = dto.RegisterStudentRequestBody{
			Fullname: "Vincent L Hubbard",
			Email:    "vincentlhubbard@edu.ac.id",
			Password: "123abcABC!",
			MajorID:  &majorID,
		}
	)
	_, err := authService.RegisterByEmailAndPassword(ctx, &payload)
	if asserts.Error(err) {
		asserts.Equal(err.Error(), "error code 409")
	}
}
