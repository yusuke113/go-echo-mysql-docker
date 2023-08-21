package router

import (
	"go/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Setup(uc controller.ITaskController) *echo.Echo {
	e := echo.New()

	e.GET("/", hello)

	// ユーザー
	e.GET("/tasks", uc.GetList)

	return e
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
