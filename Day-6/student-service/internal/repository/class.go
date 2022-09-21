package repository

import (
	"context"
	"strings"

	"student-service/internal/dto"
	"student-service/internal/model"
	pkgdto "student-service/pkg/dto"

	"gorm.io/gorm"
)

type Class interface {
	FindAll(ctx context.Context, payload *pkgdto.SearchGetRequest, p *pkgdto.Pagination) ([]model.Class, *pkgdto.PaginationInfo, error)
	FindByID(ctx context.Context, id uint) (model.Class, error)
	Save(ctx context.Context, class *dto.CreateClassRequestBody) (model.Class, error)
	Edit(ctx context.Context, oldclass *model.Class, updateData *dto.UpdateClassRequestBody) (*model.Class, error)
	Destroy(ctx context.Context, class *model.Class) (*model.Class, error)
	ExistByName(ctx context.Context, name string) (bool, error)
}

type class struct {
	Db *gorm.DB
}

func NewClassRepository(db *gorm.DB) *class {
	return &class{
		db,
	}
}

func (r *class) FindAll(ctx context.Context, payload *pkgdto.SearchGetRequest, pagination *pkgdto.Pagination) ([]model.Class, *pkgdto.PaginationInfo, error) {
	var classes []model.Class
	var count int64

	query := r.Db.WithContext(ctx).Model(&model.Class{})

	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("lower(name) LIKE ?", search)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := pkgdto.GetLimitOffset(pagination)

	err := query.Limit(limit).Offset(offset).Find(&classes).Error

	return classes, pkgdto.CheckInfoPagination(pagination, count), err
}

func (r *class) FindByID(ctx context.Context, id uint) (model.Class, error) {
	var class model.Class
	if err := r.Db.WithContext(ctx).Model(&model.Class{}).Where("id = ?", id).First(&class).Error; err != nil {
		return class, err
	}
	return class, nil
}

func (r *class) Save(ctx context.Context, class *dto.CreateClassRequestBody) (model.Class, error) {
	newClass := model.Class{
		Name: *class.Name,
	}
	if err := r.Db.WithContext(ctx).Save(&newClass).Error; err != nil {
		return newClass, err
	}
	return newClass, nil
}

func (r *class) Edit(ctx context.Context, oldClass *model.Class, updateData *dto.UpdateClassRequestBody) (*model.Class, error) {
	if updateData.Name != nil {
		oldClass.Name = *updateData.Name
	}

	if err := r.Db.WithContext(ctx).Save(oldClass).Find(oldClass).Error; err != nil {
		return nil, err
	}

	return oldClass, nil
}

func (r *class) Destroy(ctx context.Context, class *model.Class) (*model.Class, error) {
	if err := r.Db.WithContext(ctx).Delete(class).Error; err != nil {
		return nil, err
	}
	return class, nil
}

func (r *class) ExistByName(ctx context.Context, name string) (bool, error) {
	var (
		count   int64
		isExist bool
	)
	if err := r.Db.WithContext(ctx).Model(&model.Class{}).Where("name = ?", name).Count(&count).Error; err != nil {
		return isExist, err
	}
	if count > 0 {
		isExist = true
	}
	return isExist, nil
}
