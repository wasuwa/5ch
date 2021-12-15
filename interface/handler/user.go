package handler

import (
	"bbs/usecase"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	Post() echo.HandlerFunc
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(uu usecase.UserUsecase) UserHandler {
	return &userHandler{userUsecase: uu}
}

func (th *userHandler) Post() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

func RoutingUsers(e *echo.Echo, uh UserHandler) {
	e.POST("/users", uh.Post())
}
