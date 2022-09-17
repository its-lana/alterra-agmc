package controllers

import (
	"Task-Code/config"
	"Task-Code/helper"
	"Task-Code/lib/database"
	"Task-Code/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllUser(c echo.Context) error {
	users, e := database.GetAllUsers()
	if e != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, e.Error())
	}

	return helper.WrapResponse(http.StatusOK, "Success Get Users", &users).WriteToResponseBody(c.Response())

}

func GetUserByID(c echo.Context) error {
	id := c.Param("id")

	user, e := database.GetUserByID(id)

	if e != nil {
		return helper.WrapResponse(http.StatusBadRequest, "User not found!", &models.User{}).WriteToResponseBody(c.Response())
	}

	return helper.WrapResponse(http.StatusOK, "Success Get User By Id", &user).WriteToResponseBody(c.Response())
}

func AddNewUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	if err := user.ValidatorSanitizer(); err != nil {
		return helper.WrapResponse(http.StatusBadRequest, err.Error(), &models.User{}).WriteToResponseBody(c.Response())
	}

	if err := config.DB.Save(&user).Error; err != nil {
		return helper.WrapResponse(http.StatusBadRequest, "Failed Add New User!", &models.User{}).WriteToResponseBody(c.Response())
	}
	return helper.WrapResponse(http.StatusOK, "New User Added Successfully!", &user).WriteToResponseBody(c.Response())
}

func UpdateUser(c echo.Context) error {

	idParams := c.Param("id")
	id, _ := strconv.Atoi(idParams)

	user := models.User{}
	c.Bind(&user)

	if rowsAff := config.DB.Model(&user).Where("id = ?", id).Updates(user).RowsAffected; rowsAff == 0 {
		return helper.WrapResponse(http.StatusBadRequest, "Update Failed! User Id not found!", &models.User{}).WriteToResponseBody(c.Response())
	}

	return helper.WrapResponse(http.StatusOK, "User Updated Successfully!", &user).WriteToResponseBody(c.Response())
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")

	deletedUser, e := database.DeleteUserByID(id)

	if e != nil {
		return helper.WrapResponse(http.StatusBadRequest, e.Error(), &models.User{}).WriteToResponseBody(c.Response())
	}
	return helper.WrapResponse(http.StatusOK, "User Deleted Successfully!", deletedUser).WriteToResponseBody(c.Response())
}

func LoginUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	users, e := database.LoginUser(&user)

	if e != nil {
		if e.Error() == "record not found" {
			return echo.NewHTTPError(http.StatusBadRequest, "wrong email or password")
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, e.Error())
		}
	}
	user.Password = ""
	return helper.WrapResponse(http.StatusOK, "login successfully!", &users).WriteToResponseBody(c.Response())

}
