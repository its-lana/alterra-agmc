package dto

import (
	"time"

	"gorm.io/gorm"
)

type (
	CreateClassRequestBody struct {
		Name *string `json:"name" validate:"required"`
	}
	UpdateClassRequestBody struct {
		ID   *uint   `param:"id" validate:"required"`
		Name *string `json:"name" validate:"required"`
	}
	ClassResponse struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}
	ClassWithCUDResponse struct {
		ClassResponse
		CreatedAt time.Time       `json:"created_at"`
		UpdatedAt time.Time       `json:"updated_at"`
		DeletedAt *gorm.DeletedAt `json:"deleted_at"`
	}
)
