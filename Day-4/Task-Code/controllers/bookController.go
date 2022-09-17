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
		return echo.NewHTTPError(http.StatusInternalServerError, e.Error())
	}

	return helper.WrapResponse(http.StatusOK, "Success Get Books", &books).WriteToResponseBody(c.Response())

}

func GetBookByID(c echo.Context) error {
	id := c.Param("id")

	book, e := database.GetBookByID(id)

	if e != nil {
		return helper.WrapResponse(http.StatusBadRequest, "Book not found!", &models.Book{}).WriteToResponseBody(c.Response())
	}

	return helper.WrapResponse(http.StatusOK, "Success Get Book By Id", &book).WriteToResponseBody(c.Response())
}

func AddNewBook(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)

	if err := book.ValidatorSanitizer(); err != nil {
		return helper.WrapResponse(http.StatusBadRequest, err.Error(), &models.User{}).WriteToResponseBody(c.Response())
	}

	if err := config.DB.Save(&book).Error; err != nil {
		return helper.WrapResponse(http.StatusBadRequest, "Failed Add New Book!", &models.Book{}).WriteToResponseBody(c.Response())
	}
	return helper.WrapResponse(http.StatusOK, "New Book Added Successfully!", &book).WriteToResponseBody(c.Response())
}

func UpdateBook(c echo.Context) error {

	id := c.Param("id")

	book := models.Book{}
	c.Bind(&book)

	if rowsAff := config.DB.Model(&book).Where("id = ?", id).Updates(book).RowsAffected; rowsAff == 0 {
		return helper.WrapResponse(http.StatusBadRequest, "Update Failed! Book Id not found!", &models.Book{}).WriteToResponseBody(c.Response())
	}

	return helper.WrapResponse(http.StatusOK, "Book Updated Successfully!", &book).WriteToResponseBody(c.Response())
}

func DeleteBook(c echo.Context) error {
	id := c.Param("id")

	_, e := database.DeleteBookByID(id)

	if e != nil {
		return helper.WrapResponse(http.StatusBadRequest, e.Error(), &models.Book{}).WriteToResponseBody(c.Response())
	}
	return helper.WrapResponse(http.StatusOK, "Book Deleted Successfully!", &models.Book{}).WriteToResponseBody(c.Response())
}
