package dto

import (
	"time"

	"gorm.io/gorm"
)

type (
	CreateMajorRequestBody struct {
		Name *string `json:"name" validate:"required"`
	}
	UpdateMajorRequestBody struct {
		ID   *uint   `param:"id" validate:"required"`
		Name *string `json:"name" validate:"required"`
	}
	MajorResponse struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}
	MajorWithCUDResponse struct {
		MajorResponse
		CreatedAt time.Time       `json:"created_at"`
		UpdatedAt time.Time       `json:"updated_at"`
		DeletedAt *gorm.DeletedAt `json:"deleted_at"`
	}
)
