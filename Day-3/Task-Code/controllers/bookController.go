package controllers

import (
	"Task-Code/config"
	"Task-Code/helper"
	"Task-Code/lib/database"
	"Task-Code/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllBooks(c echo.Context) error {
	books, e := database.GetAllBooks()
	if e != nil {
		return helper.WrapResponse(http.StatusBadRequest, "Failed Get Books", &models.Book{}).WriteToResponseBody(c.Response())
	}

	return helper.WrapResponse(http.StatusOK, "Success Get Books", &books).WriteToResponseBody(c.Response())

}

func GetBookByID(c echo.Context) error {
	id := c.Param("id")

	book, e := database.GetBookByID(id)

	if e != nil {
		return helper.WrapResponse(http.StatusBadRequest, "Failed Get Book by Id", &models.Book{}).WriteToResponseBody(c.Response())
	}

	return helper.WrapResponse(http.StatusOK, "Success Get Book By Id", &book).WriteToResponseBody(c.Response())
}

func AddNewBook(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)

	if err := config.DB.Save(&book).Error; err != nil {
		return helper.WrapResponse(http.StatusBadRequest, "Failed Add New Book!", &models.Book{}).WriteToResponseBody(c.Response())
	}
	return helper.WrapResponse(http.StatusOK, "New Book Added Successfully!", &book).WriteToResponseBody(c.Response())
}

func UpdateBook(c echo.Context) error {

	id := c.Param("id")

	book := models.Book{}
	c.Bind(&book)

	if err := config.DB.Model(&book).Where("id = ?", id).Updates(book).Error; err != nil {
		return helper.WrapResponse(http.StatusBadRequest, "Failed Update Book!", &models.Book{}).WriteToResponseBody(c.Response())
	}

	return helper.WrapResponse(http.StatusOK, "Book Updated Successfully!", &book).WriteToResponseBody(c.Response())
}

func DeleteBook(c echo.Context) error {
	id := c.Param("id")

	_, e := database.DeleteBookByID(id)

	if e != nil {
		return helper.WrapResponse(http.StatusBadRequest, "Failed Delete Book!", &models.Book{}).WriteToResponseBody(c.Response())
	}
	return helper.WrapResponse(http.StatusOK, "Book Deleted Successfully!", &models.Book{}).WriteToResponseBody(c.Response())
}
