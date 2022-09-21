package http

import (
	"student-service/internal/app/auth"
	"student-service/internal/app/class"
	"student-service/internal/app/major"
	"student-service/internal/app/student"
	"student-service/internal/factory"
	"student-service/pkg/util"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {
	e.Validator = &util.CustomValidator{Validator: validator.New()}

	e.GET("/status", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "OK"})
	})
	v1 := e.Group("/api/v1")
	student.NewHandler(f).Route(v1.Group("/students"))
	auth.NewHandler(f).Route(v1.Group("/auth"))
	major.NewHandler(f).Route(v1.Group("/majors"))
	class.NewHandler(f).Route(v1.Group("/classes"))
}
