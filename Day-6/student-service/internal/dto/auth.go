package dto

import "github.com/golang-jwt/jwt/v4"

type (
	RegisterStudentRequestBody struct {
		Fullname string `json:"fullname" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
		ClassID  *uint  `json:"class_id"`
		MajorID  *uint  `json:"major_id" validate:"required"`
	}

	ByEmailAndPasswordRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	JWTClaims struct {
		BID     uint   `json:"user_id"`
		Email   string `json:"email"`
		ClassID uint   `json:"class_id"`
		MajorID uint   `json:"major_id"`
		jwt.RegisteredClaims
	}
)

func (r *RegisterStudentRequestBody) FillDefaults() {
	var defaultClassID uint = 1
	if r.ClassID == nil {
		r.ClassID = &defaultClassID
	}
}
