package middlewares

import (
	"Task-Code/helper"
	"Task-Code/models"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func UserAuthMiddlewares() func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return helper.WrapResponse(http.StatusUnauthorized, "You are not Authorized!", &models.User{}).WriteToResponseBody(c.Response())
			}

			token := strings.Split(authHeader, " ")[1]
			userId, e := ValidateToken(token)
			if userId == 0 || e != nil {
				return helper.WrapResponse(http.StatusUnauthorized, "You are not Authorized!", &models.User{}).WriteToResponseBody(c.Response())
			}

			c.Set("userId", userId)
			return next(c)

		}
	}
}
