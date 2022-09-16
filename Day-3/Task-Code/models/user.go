package models

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Token    string `json:"token" form:"token"`
}

func (user *User) ValidatorSanitizer() error {
	if user.Name == "" {
		return fmt.Errorf("name field is required")
	}
	if user.Email == "" {
		return fmt.Errorf("email field is required")
	}
	if user.Password == "" {
		return fmt.Errorf("password field is required")
	}
	return nil
}
