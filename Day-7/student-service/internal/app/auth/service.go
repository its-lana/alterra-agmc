package auth

import (
	"context"
	"errors"

	"student-service/internal/dto"
	"student-service/internal/factory"
	"student-service/internal/pkg/util"
	"student-service/internal/repository"
	"student-service/pkg/constant"
	pkgutil "student-service/pkg/util"
	res "student-service/pkg/util/response"
)

type service struct {
	StudentRepository repository.Student
}

type Service interface {
	LoginByEmailAndPassword(ctx context.Context, payload *dto.ByEmailAndPasswordRequest) (*dto.StudentWithJWTResponse, error)
	RegisterByEmailAndPassword(ctx context.Context, payload *dto.RegisterStudentRequestBody) (*dto.StudentWithJWTResponse, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		StudentRepository: f.StudentRepository,
	}
}

func (s *service) LoginByEmailAndPassword(ctx context.Context, payload *dto.ByEmailAndPasswordRequest) (*dto.StudentWithJWTResponse, error) {
	var result *dto.StudentWithJWTResponse

	data, err := s.StudentRepository.FindByEmail(ctx, &payload.Email)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return result, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	if !(pkgutil.CompareHashPassword(payload.Password, data.Password)) {
		return result, res.ErrorBuilder(
			&res.ErrorConstant.EmailOrPasswordIncorrect,
			errors.New(res.ErrorConstant.EmailOrPasswordIncorrect.Response.Meta.Message),
		)
	}

	claims := util.CreateJWTClaims(data.Email, data.ID, data.ClassID, data.MajorID)
	token, err := util.CreateJWTToken(claims)
	if err != nil {
		return result, res.ErrorBuilder(
			&res.ErrorConstant.InternalServerError,
			errors.New("error when generating token"),
		)
	}

	result = &dto.StudentWithJWTResponse{
		StudentResponse: dto.StudentResponse{
			ID:       data.ID,
			Fullname: data.Fullname,
			Email:    data.Email,
		},
		JWT: token,
	}

	return result, nil
}

func (s *service) RegisterByEmailAndPassword(ctx context.Context, payload *dto.RegisterStudentRequestBody) (*dto.StudentWithJWTResponse, error) {
	var result *dto.StudentWithJWTResponse
	isExist, err := s.StudentRepository.ExistByEmail(ctx, &payload.Email)
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	if isExist {
		return result, res.ErrorBuilder(&res.ErrorConstant.Duplicate, errors.New("student already exists"))
	}

	hashedPassword, err := pkgutil.HashPassword(payload.Password)
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	payload.Password = hashedPassword

	data, err := s.StudentRepository.Save(ctx, payload)
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	claims := util.CreateJWTClaims(data.Email, data.ID, data.ClassID, data.MajorID)
	token, err := util.CreateJWTToken(claims)
	if err != nil {
		return result, res.ErrorBuilder(
			&res.ErrorConstant.InternalServerError,
			errors.New("error when generating token"),
		)
	}

	result = &dto.StudentWithJWTResponse{
		StudentResponse: dto.StudentResponse{
			ID:       data.ID,
			Fullname: data.Fullname,
			Email:    data.Email,
		},
		JWT: token,
	}

	return result, nil
}
