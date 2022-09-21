package repository

import (
	"context"
	"strings"

	"student-service/internal/dto"
	"student-service/internal/model"
	pkgdto "student-service/pkg/dto"
	"student-service/pkg/util"

	"gorm.io/gorm"
)

type Student interface {
	FindAll(ctx context.Context, payload *pkgdto.SearchGetRequest, p *pkgdto.Pagination) ([]model.Student, *pkgdto.PaginationInfo, error)
	FindByID(ctx context.Context, id uint, usePreload bool) (model.Student, error)
	FindByEmail(ctx context.Context, email *string) (*model.Student, error)
	ExistByEmail(ctx context.Context, email *string) (bool, error)
	ExistByID(ctx context.Context, id uint) (bool, error)
	Save(ctx context.Context, student *dto.RegisterStudentRequestBody) (model.Student, error)
	Edit(ctx context.Context, oldStudent *model.Student, updateData *dto.UpdateStudentRequestBody) (*model.Student, error)
	Destroy(ctx context.Context, student *model.Student) (*model.Student, error)
}

type student struct {
	Db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) *student {
	return &student{
		db,
	}
}

func (r *student) FindAll(ctx context.Context, payload *pkgdto.SearchGetRequest, pagination *pkgdto.Pagination) ([]model.Student, *pkgdto.PaginationInfo, error) {
	var users []model.Student
	var count int64

	query := r.Db.WithContext(ctx).Model(&model.Student{})

	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("lower(fullname) LIKE ? or lower(email) Like ? ", search, search)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := pkgdto.GetLimitOffset(pagination)

	err := query.Limit(limit).Offset(offset).Find(&users).Error

	return users, pkgdto.CheckInfoPagination(pagination, count), err
}

func (r *student) FindByID(ctx context.Context, id uint, usePreload bool) (model.Student, error) {
	var user model.Student
	q := r.Db.WithContext(ctx).Model(&model.Student{}).Where("id = ?", id)
	if usePreload {
		q = q.Preload("Major").Preload("Class")
	}
	err := q.First(&user).Error
	return user, err
}

func (r *student) FindByEmail(ctx context.Context, email *string) (*model.Student, error) {
	var data model.Student
	err := r.Db.WithContext(ctx).Where("email = ?", email).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *student) ExistByEmail(ctx context.Context, email *string) (bool, error) {
	var (
		count   int64
		isExist bool
	)
	if err := r.Db.WithContext(ctx).Model(&model.Student{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return isExist, err
	}
	if count > 0 {
		isExist = true
	}
	return isExist, nil
}

func (r *student) ExistByID(ctx context.Context, id uint) (bool, error) {
	var (
		count   int64
		isExist bool
	)
	if err := r.Db.WithContext(ctx).Model(&model.Student{}).Where("id = ?", id).Count(&count).Error; err != nil {
		return isExist, err
	}
	if count > 0 {
		isExist = true
	}
	return isExist, nil
}

func (r *student) Save(ctx context.Context, student *dto.RegisterStudentRequestBody) (model.Student, error) {
	newStudent := model.Student{
		Fullname: student.Fullname,
		Email:    student.Email,
		Password: student.Password,
		ClassID:  *student.ClassID,
		MajorID:  *student.MajorID,
	}
	if err := r.Db.WithContext(ctx).Save(&newStudent).Error; err != nil {
		return newStudent, err
	}
	return newStudent, nil
}

func (r *student) Edit(ctx context.Context, oldStudent *model.Student, updateData *dto.UpdateStudentRequestBody) (*model.Student, error) {
	if updateData.Fullname != nil {
		oldStudent.Fullname = *updateData.Fullname
	}
	if updateData.Email != nil {
		oldStudent.Email = *updateData.Email
	}
	if updateData.Password != nil {
		hashedPassword, err := util.HashPassword(*updateData.Password)
		if err != nil {
			return nil, err
		}
		oldStudent.Password = hashedPassword
	}
	if updateData.MajorID != nil {
		oldStudent.MajorID = *updateData.MajorID
	}
	if updateData.ClassID != nil {
		oldStudent.ClassID = *updateData.ClassID
	}

	if err := r.Db.
		WithContext(ctx).
		Save(oldStudent).
		Preload("Major").
		Preload("Class").
		Find(oldStudent).
		Error; err != nil {
		return nil, err
	}

	return oldStudent, nil
}

func (r *student) Destroy(ctx context.Context, student *model.Student) (*model.Student, error) {
	if err := r.Db.WithContext(ctx).Delete(student).Error; err != nil {
		return nil, err
	}
	return student, nil
}
