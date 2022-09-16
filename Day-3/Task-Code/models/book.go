package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title     string `json:"title" form:"title"`
	Author    string `json:"author" form:"author"`
	TotalPage uint   `json:"total_page" form:"total_page"`
}

func (book *Book) ValidatorSanitizer() error {
	if book.Title == "" {
		return fmt.Errorf("title field is required")
	}
	if book.Author == "" {
		return fmt.Errorf("author field is required")
	}
	return nil
}
