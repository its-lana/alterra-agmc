package major

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
	MajorRepository repository.Major
}

type Service interface {
	Find(ctx context.Context, payload *pkgdto.SearchGetRequest) (*pkgdto.SearchGetResponse[dto.MajorResponse], error)
	FindByID(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.MajorResponse, error)
	Store(ctx context.Context, payload *dto.CreateMajorRequestBody) (*dto.MajorResponse, error)
	UpdateById(ctx context.Context, payload *dto.UpdateMajorRequestBody) (*dto.MajorResponse, error)
	DeleteById(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.MajorWithCUDResponse, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		MajorRepository: f.MajorRepository,
	}
}

func (s *service) Find(ctx context.Context, payload *pkgdto.SearchGetRequest) (*pkgdto.SearchGetResponse[dto.MajorResponse], error) {
	majors, info, err := s.MajorRepository.FindAll(ctx, payload, &payload.Pagination)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	var data []dto.MajorResponse

	for _, major := range majors {
		data = append(data, dto.MajorResponse{
			ID:   major.ID,
			Name: major.Name,
		})

	}

	result := new(pkgdto.SearchGetResponse[dto.MajorResponse])
	result.Data = data
	result.PaginationInfo = *info

	return result, nil
}
func (s *service) FindByID(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.MajorResponse, error) {
	var result dto.MajorResponse
	data, err := s.MajorRepository.FindByID(ctx, payload.ID)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return &dto.MajorResponse{}, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return &dto.MajorResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result.ID = data.ID
	result.Name = data.Name

	return &result, nil
}

func (s *service) Store(ctx context.Context, payload *dto.CreateMajorRequestBody) (*dto.MajorResponse, error) {
	var result dto.MajorResponse
	isExist, err := s.MajorRepository.ExistByName(ctx, *payload.Name)
	if err != nil {
		return &result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	if isExist {
		return &result, res.ErrorBuilder(&res.ErrorConstant.Duplicate, errors.New("major already exists"))
	}

	data, err := s.MajorRepository.Save(ctx, payload)
	if err != nil {
		return &result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result.ID = data.ID
	result.Name = data.Name

	return &result, nil
}

func (s *service) UpdateById(ctx context.Context, payload *dto.UpdateMajorRequestBody) (*dto.MajorResponse, error) {
	major, err := s.MajorRepository.FindByID(ctx, *payload.ID)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return &dto.MajorResponse{}, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return &dto.MajorResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	_, err = s.MajorRepository.Edit(ctx, &major, payload)
	if err != nil {
		return &dto.MajorResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	var result dto.MajorResponse
	result.ID = major.ID
	result.Name = major.Name

	return &result, nil
}
func (s *service) DeleteById(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.MajorWithCUDResponse, error) {
	major, err := s.MajorRepository.FindByID(ctx, payload.ID)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return &dto.MajorWithCUDResponse{}, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return &dto.MajorWithCUDResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	_, err = s.MajorRepository.Destroy(ctx, &major)
	if err != nil {
		return &dto.MajorWithCUDResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := &dto.MajorWithCUDResponse{
		MajorResponse: dto.MajorResponse{
			ID:   major.ID,
			Name: major.Name,
		},
		CreatedAt: major.CreatedAt,
		UpdatedAt: major.UpdatedAt,
		DeletedAt: major.DeletedAt,
	}

	return result, nil
}
