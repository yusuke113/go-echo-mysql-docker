package controller

import (
	"go/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ITaskController interface {
	GetList(c echo.Context) error
}

type taskController struct {
	uu usecase.ITaskUseCase
}

func NewTaskController(uu usecase.ITaskUseCase) ITaskController {
	return &taskController{uu}
}

func (uc *taskController) GetList(c echo.Context) error {
	tasksRes, err := uc.uu.GetList()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, tasksRes)
}
