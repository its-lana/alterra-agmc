package routes

import (
	"Dynamic-CRUD/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	v1 := e.Group("/v1")
	users := v1.Group("/users")

	users.GET("", controllers.GetAllUser)
	users.GET("/:id", controllers.GetUserByID)
	users.POST("", controllers.AddNewUser)
	users.PUT("/:id", controllers.UpdateUser)
	users.DELETE("/:id", controllers.DeleteUser)

	return e
}
