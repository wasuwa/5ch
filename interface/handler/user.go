package handler

import (
	"bbs/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	Post() echo.HandlerFunc
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

type requestUser struct {
	Name     string `json:"name"`
	Email    string `json:"email`
	Password string `json:"password"`
}

func NewUserHandler(uu usecase.UserUsecase) UserHandler {
	return &userHandler{userUsecase: uu}
}

func (uh *userHandler) Post() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(requestUser)
		if err := c.Bind(req); err != nil {
			return c.JSONPretty(http.StatusBadRequest, err.Error(), " ")
		}
		u, err := uh.userUsecase.Create(req.Name, req.Email, req.Password)
		if err != nil {
			return c.JSONPretty(http.StatusBadRequest, err.Error(), " ")
		}
		return c.JSONPretty(http.StatusCreated, u, " ")
	}
}

func RoutingUsers(e *echo.Echo, uh UserHandler) {
	e.POST("/users", uh.Post())
}
