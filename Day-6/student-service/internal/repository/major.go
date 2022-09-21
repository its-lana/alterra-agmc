package repository

import (
	"context"
	"strings"

	"student-service/internal/dto"
	"student-service/internal/model"
	pkgdto "student-service/pkg/dto"

	"gorm.io/gorm"
)

type Major interface {
	FindAll(ctx context.Context, payload *pkgdto.SearchGetRequest, pagination *pkgdto.Pagination) ([]model.Major, *pkgdto.PaginationInfo, error)
	FindByID(ctx context.Context, id uint) (model.Major, error)
	Save(ctx context.Context, major *dto.CreateMajorRequestBody) (model.Major, error)
	Edit(ctx context.Context, oldStudent *model.Major, updateData *dto.UpdateMajorRequestBody) (*model.Major, error)
	Destroy(ctx context.Context, major *model.Major) (*model.Major, error)
	ExistByName(ctx context.Context, name string) (bool, error)
}

type major struct {
	Db *gorm.DB
}

func NewMajorRepository(db *gorm.DB) *major {
	return &major{
		db,
	}
}

func (r *major) FindAll(ctx context.Context, payload *pkgdto.SearchGetRequest, pagination *pkgdto.Pagination) ([]model.Major, *pkgdto.PaginationInfo, error) {
	var majors []model.Major
	var count int64

	query := r.Db.WithContext(ctx).Model(&model.Major{})

	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("lower(name) LIKE ?", search)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := pkgdto.GetLimitOffset(pagination)

	err := query.Limit(limit).Offset(offset).Find(&majors).Error

	return majors, pkgdto.CheckInfoPagination(pagination, count), err
}

func (r *major) FindByID(ctx context.Context, id uint) (model.Major, error) {
	var major model.Major
	if err := r.Db.WithContext(ctx).Model(&model.Major{}).Where("id = ?", id).First(&major).Error; err != nil {
		return major, err
	}
	return major, nil
}

func (r *major) Save(ctx context.Context, major *dto.CreateMajorRequestBody) (model.Major, error) {
	newMajor := model.Major{
		Name: *major.Name,
	}
	if err := r.Db.WithContext(ctx).Save(&newMajor).Error; err != nil {
		return newMajor, err
	}
	return newMajor, nil
}

func (r *major) Edit(ctx context.Context, oldMajor *model.Major, updateData *dto.UpdateMajorRequestBody) (*model.Major, error) {
	if updateData.Name != nil {
		oldMajor.Name = *updateData.Name
	}

	if err := r.Db.WithContext(ctx).Save(oldMajor).Find(oldMajor).Error; err != nil {
		return nil, err
	}

	return oldMajor, nil
}

func (r *major) Destroy(ctx context.Context, major *model.Major) (*model.Major, error) {
	if err := r.Db.WithContext(ctx).Delete(major).Error; err != nil {
		return nil, err
	}
	return major, nil
}

func (r *major) ExistByName(ctx context.Context, name string) (bool, error) {
	var (
		count   int64
		isExist bool
	)
	if err := r.Db.WithContext(ctx).Model(&model.Major{}).Where("name = ?", name).Count(&count).Error; err != nil {
		return isExist, err
	}
	if count > 0 {
		isExist = true
	}
	return isExist, nil
}
