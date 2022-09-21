package class

import (
	"context"
	"errors"

	"student-service/internal/dto"
	"student-service/internal/factory"
	"student-service/internal/repository"
	"student-service/pkg/constant"
	pkgdto "student-service/pkg/dto"
	res "student-service/pkg/util/response"
)

type service struct {
	ClassRepository repository.Class
}

type Service interface {
	Find(ctx context.Context, payload *pkgdto.SearchGetRequest) (*pkgdto.SearchGetResponse[dto.ClassResponse], error)
	FindByID(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.ClassResponse, error)
	Store(ctx context.Context, payload *dto.CreateClassRequestBody) (*dto.ClassResponse, error)
	UpdateById(ctx context.Context, payload *dto.UpdateClassRequestBody) (*dto.ClassResponse, error)
	DeleteById(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.ClassWithCUDResponse, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		ClassRepository: f.ClassRepository,
	}
}

func (s *service) Find(ctx context.Context, payload *pkgdto.SearchGetRequest) (*pkgdto.SearchGetResponse[dto.ClassResponse], error) {
	classes, info, err := s.ClassRepository.FindAll(ctx, payload, &payload.Pagination)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	var data []dto.ClassResponse

	for _, class := range classes {
		data = append(data, dto.ClassResponse{
			ID:   class.ID,
			Name: class.Name,
		})

	}

	result := new(pkgdto.SearchGetResponse[dto.ClassResponse])
	result.Data = data
	result.PaginationInfo = *info

	return result, nil
}
func (s *service) FindByID(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.ClassResponse, error) {
	var result dto.ClassResponse
	data, err := s.ClassRepository.FindByID(ctx, payload.ID)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return &dto.ClassResponse{}, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return &dto.ClassResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result.ID = data.ID
	result.Name = data.Name

	return &result, nil
}

func (s *service) Store(ctx context.Context, payload *dto.CreateClassRequestBody) (*dto.ClassResponse, error) {
	var result dto.ClassResponse
	isExist, err := s.ClassRepository.ExistByName(ctx, *payload.Name)
	if err != nil {
		return &result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	if isExist {
		return &result, res.ErrorBuilder(&res.ErrorConstant.Duplicate, errors.New("class already exists"))
	}

	data, err := s.ClassRepository.Save(ctx, payload)
	if err != nil {
		return &result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result.ID = data.ID
	result.Name = data.Name

	return &result, nil
}

func (s *service) UpdateById(ctx context.Context, payload *dto.UpdateClassRequestBody) (*dto.ClassResponse, error) {
	class, err := s.ClassRepository.FindByID(ctx, *payload.ID)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return &dto.ClassResponse{}, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return &dto.ClassResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	_, err = s.ClassRepository.Edit(ctx, &class, payload)
	if err != nil {
		return &dto.ClassResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	var result dto.ClassResponse
	result.ID = class.ID
	result.Name = class.Name

	return &result, nil
}
func (s *service) DeleteById(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.ClassWithCUDResponse, error) {
	class, err := s.ClassRepository.FindByID(ctx, payload.ID)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return &dto.ClassWithCUDResponse{}, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return &dto.ClassWithCUDResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	_, err = s.ClassRepository.Destroy(ctx, &class)
	if err != nil {
		return &dto.ClassWithCUDResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := &dto.ClassWithCUDResponse{
		ClassResponse: dto.ClassResponse{
			ID:   class.ID,
			Name: class.Name,
		},
		CreatedAt: class.CreatedAt,
		UpdatedAt: class.UpdatedAt,
		DeletedAt: class.DeletedAt,
	}

	return result, nil
}
