package seeder

import (
	"Task-Code/config"
	"Task-Code/models"
	"log"

	"gorm.io/gorm"
)

type seed struct {
	DB *gorm.DB
}

func NewSeeder() *seed {
	config.InitDB()
	return &seed{DB: config.DB}
}

func (s *seed) UserSeed() {
	users := []models.User{
		{
			Model: gorm.Model{
				ID: 1,
			},
			Name:     "test1",
			Email:    "test1@mail.com",
			Password: "1234",
		},
		{
			Model: gorm.Model{
				ID: 2,
			},
			Name:     "test2",
			Email:    "test2@mail.com",
			Password: "1234",
		},
	}
	if err := s.DB.Create(&users).Error; err != nil {
		log.Printf("cannot seed data users, error : %v\n", err)
	}
	log.Println("success seed data users")
}

func (s *seed) BookSeed() {
	books := []models.Book{
		{
			Model: gorm.Model{
				ID: 1,
			},
			Title:     "Test Book 1",
			Author:    "Tester 1",
			TotalPage: 340,
		},
		{
			Model: gorm.Model{
				ID: 2,
			},
			Title:     "Test Book 2",
			Author:    "Tester 2",
			TotalPage: 420,
		},
	}
	if err := s.DB.Create(&books).Error; err != nil {
		log.Printf("cannot seed data books, error : %v\n", err)
	}
	log.Println("success seed data books")
}

func (s *seed) UserDelete() {
	s.DB.Exec("DELETE FROM users")
}

func (s *seed) BookDelete() {
	s.DB.Exec("DELETE FROM books")
}
