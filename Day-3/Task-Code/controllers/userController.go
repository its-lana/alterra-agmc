package controllers

import (
	"Task-Code/config"
	"Task-Code/helper"
	"Task-Code/lib/database"
	"Task-Code/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllUser(c echo.Context) error {
	users, e := database.GetAllUsers()
	if e != nil {
		return helper.WrapResponse(http.StatusBadRequest, "Failed Get Users", &models.User{}).WriteToResponseBody(c.Response())
	}

	return helper.WrapResponse(http.StatusOK, "Success Get Users", &users).WriteToResponseBody(c.Response())

}

func GetUserByID(c echo.Context) error {
	id := c.Param("id")

	user, e := database.GetUserByID(id)

	if e != nil {
		return helper.WrapResponse(http.StatusBadRequest, "Failed Get User by Id", &models.User{}).WriteToResponseBody(c.Response())
	}

	return helper.WrapResponse(http.StatusOK, "Success Get User By Id", &user).WriteToResponseBody(c.Response())
}

func AddNewUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	if err := config.DB.Save(&user).Error; err != nil {
		return helper.WrapResponse(http.StatusBadRequest, "Failed Add New User!", &models.User{}).WriteToResponseBody(c.Response())
	}
	return helper.WrapResponse(http.StatusOK, "New User Added Successfully!", &user).WriteToResponseBody(c.Response())
}

func UpdateUser(c echo.Context) error {

	id := c.Param("id")

	user := models.User{}
	c.Bind(&user)

	if err := config.DB.Model(&user).Where("id = ?", id).Updates(user).Error; err != nil {
		return helper.WrapResponse(http.StatusBadRequest, "Failed Update User!", &models.User{}).WriteToResponseBody(c.Response())
	}

	return helper.WrapResponse(http.StatusOK, "User Updated Successfully!", &user).WriteToResponseBody(c.Response())
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")

	_, e := database.DeleteUserByID(id)

	if e != nil {
		return helper.WrapResponse(http.StatusBadRequest, "Failed Delete User!", &models.User{}).WriteToResponseBody(c.Response())
	}
	return helper.WrapResponse(http.StatusOK, "User Deleted Successfully!", &models.User{}).WriteToResponseBody(c.Response())
}

func LoginUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	users, e := database.LoginUser(&user)

	if e != nil {
		if e.Error() == "record not found" {
			return helper.WrapResponse(http.StatusBadRequest, "wrong email or password", &models.User{}).WriteToResponseBody(c.Response())
		} else {
			return helper.WrapResponse(http.StatusInternalServerError, "internal server error", &models.User{}).WriteToResponseBody(c.Response())
		}
	}
	user.Password = ""
	return helper.WrapResponse(http.StatusOK, "login successfully!", &users).WriteToResponseBody(c.Response())

}
