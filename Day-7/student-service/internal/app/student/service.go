package student

import (
	"context"

	"student-service/internal/dto"
	"student-service/internal/factory"
	"student-service/internal/repository"
	"student-service/pkg/constant"
	pkgdto "student-service/pkg/dto"
	res "student-service/pkg/util/response"
)

type service struct {
	StudentRepository repository.Student
}

type Service interface {
	Find(ctx context.Context, payload *pkgdto.SearchGetRequest) (*pkgdto.SearchGetResponse[dto.StudentResponse], error)
	FindByID(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.StudentDetailResponse, error)
	UpdateById(ctx context.Context, payload *dto.UpdateStudentRequestBody) (*dto.StudentDetailResponse, error)
	DeleteById(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.StudentWithCUDResponse, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		StudentRepository: f.StudentRepository,
	}
}

func (s *service) Find(ctx context.Context, payload *pkgdto.SearchGetRequest) (*pkgdto.SearchGetResponse[dto.StudentResponse], error) {
	students, info, err := s.StudentRepository.FindAll(ctx, payload, &payload.Pagination)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	var data []dto.StudentResponse

	for _, student := range students {
		data = append(data, dto.StudentResponse{
			ID:       student.ID,
			Fullname: student.Fullname,
			Email:    student.Email,
		})

	}

	result := new(pkgdto.SearchGetResponse[dto.StudentResponse])
	result.Data = data
	result.PaginationInfo = *info

	return result, nil
}

func (s *service) FindByID(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.StudentDetailResponse, error) {
	data, err := s.StudentRepository.FindByID(ctx, payload.ID, true)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return &dto.StudentDetailResponse{}, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return &dto.StudentDetailResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := &dto.StudentDetailResponse{
		StudentResponse: dto.StudentResponse{
			ID:       data.ID,
			Fullname: data.Fullname,
			Email:    data.Email,
		},
		Class: dto.ClassResponse{
			ID:   data.Class.ID,
			Name: data.Class.Name,
		},
		Major: dto.MajorResponse{
			ID:   data.Major.ID,
			Name: data.Major.Name,
		},
	}

	return result, nil
}

func (s *service) UpdateById(ctx context.Context, payload *dto.UpdateStudentRequestBody) (*dto.StudentDetailResponse, error) {
	student, err := s.StudentRepository.FindByID(ctx, *payload.ID, false)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return &dto.StudentDetailResponse{}, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return &dto.StudentDetailResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	_, err = s.StudentRepository.Edit(ctx, &student, payload)
	if err != nil {
		return &dto.StudentDetailResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := &dto.StudentDetailResponse{
		StudentResponse: dto.StudentResponse{
			ID:       student.ID,
			Fullname: student.Fullname,
			Email:    student.Email,
		},
		Class: dto.ClassResponse{
			ID:   student.Class.ID,
			Name: student.Class.Name,
		},
		Major: dto.MajorResponse{
			ID:   student.Major.ID,
			Name: student.Major.Name,
		},
	}

	return result, nil
}

func (s *service) DeleteById(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.StudentWithCUDResponse, error) {
	student, err := s.StudentRepository.FindByID(ctx, payload.ID, false)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return &dto.StudentWithCUDResponse{}, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return &dto.StudentWithCUDResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	_, err = s.StudentRepository.Destroy(ctx, &student)
	if err != nil {
		return &dto.StudentWithCUDResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := &dto.StudentWithCUDResponse{
		StudentResponse: dto.StudentResponse{
			ID:       student.ID,
			Fullname: student.Fullname,
			Email:    student.Email,
		},
		CreatedAt: student.CreatedAt,
		UpdatedAt: student.UpdatedAt,
		DeletedAt: student.DeletedAt,
	}

	return result, nil
}
