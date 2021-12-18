package handler

import (
	"bbs/domain/model"
	"bbs/usecase"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	GetAll()  echo.HandlerFunc
	Post() echo.HandlerFunc
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

type requestUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type responseUser struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type responseAllUser struct {
	Users *[]model.User `json:"users"`
}

func NewUserHandler(uu usecase.UserUsecase) UserHandler {
	return &userHandler{userUsecase: uu}
}

func (uh *userHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := uh.userUsecase.Index()
		if err != nil {
			return c.JSONPretty(http.StatusNotFound, err.Error(), " ")
		}
		res := &responseAllUser{Users: users}
		return c.JSONPretty(http.StatusOK, res, " ")
	}
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
		res := &responseUser{
			ID:        u.GetID(),
			Name:      u.GetName(),
			Email:     u.GetEmail(),
			Password:  u.GetPassword(),
			CreatedAt: u.GetCreatedAt(),
			UpdatedAt: u.GetUpdatedAt(),
		}
		return c.JSONPretty(http.StatusCreated, res, " ")
	}
}

func RoutingUsers(e *echo.Echo, uh UserHandler) {
	e.GET("users", uh.GetAll())
	e.POST("/users", uh.Post())
}
