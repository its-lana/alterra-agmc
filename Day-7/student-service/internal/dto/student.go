package dto

import (
	"time"

	"gorm.io/gorm"
)

type (
	UpdateStudentRequestBody struct {
		ID       *uint   `param:"id" validate:"required"`
		Fullname *string `json:"fullname" validate:"omitempty"`
		Email    *string `json:"email" validate:"omitempty,email"`
		Password *string `json:"password" validate:"omitempty"`
		ClassID  *uint   `json:"class_id" validate:"omitempty"`
		MajorID  *uint   `json:"major_id" validate:"omitempty"`
	}
	StudentResponse struct {
		ID       uint   `json:"id"`
		Fullname string `json:"fullname"`
		Email    string `json:"email"`
	}
	StudentWithJWTResponse struct {
		StudentResponse
		JWT string `json:"jwt"`
	}
	StudentWithCUDResponse struct {
		StudentResponse
		CreatedAt time.Time       `json:"created_at"`
		UpdatedAt time.Time       `json:"updated_at"`
		DeletedAt *gorm.DeletedAt `json:"deleted_at"`
	}
	StudentDetailResponse struct {
		StudentResponse
		Class ClassResponse `json:"class"`
		Major MajorResponse `json:"major"`
	}
)
