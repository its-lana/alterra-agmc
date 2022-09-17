package routes

import (
	"Task-Code/controllers"
	"Task-Code/middlewares"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	middlewares.LogMiddlewares(e)

	v1 := e.Group("/v1")
	v1Auth := e.Group("/v1", middlewares.UserAuthMiddlewares())

	//user login
	v1.POST("/login", controllers.LoginUser)

	//api Book
	v1.GET("/books", controllers.GetAllBooks)
	v1.GET("/books/:id", controllers.GetBookByID)
	v1Auth.POST("/books", controllers.AddNewBook)
	v1Auth.PUT("/books/:id", controllers.UpdateBook)
	v1Auth.DELETE("/books/:id", controllers.DeleteBook)

	//api User
	v1Auth.GET("/users", controllers.GetAllUser)
	v1Auth.GET("/users/:id", controllers.GetUserByID)
	v1.POST("/users", controllers.AddNewUser)
	v1Auth.PUT("/users/:id", controllers.UpdateUser)
	v1Auth.DELETE("/users/:id", controllers.DeleteUser)

	return e
}
