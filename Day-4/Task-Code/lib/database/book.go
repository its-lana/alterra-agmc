package database

import (
	"Task-Code/config"
	"Task-Code/models"
	"errors"
)

func GetAllBooks() (interface{}, error) {
	var books []models.Book
	if e := config.DB.Find(&books).Error; e != nil {
		return nil, e
	}
	return books, nil
}

func GetBookByID(id string) (interface{}, error) {
	var book models.Book

	if e := config.DB.First(&book, id).Error; e != nil {
		return nil, e
	}
	return book, nil
}

func DeleteBookByID(id string) (interface{}, error) {
	var book models.Book

	if rowsAff := config.DB.Delete(&book, id).RowsAffected; rowsAff == 0 {
		return nil, errors.New("Delete Failed! Book Id not found!")
	}
	return book, nil
}
