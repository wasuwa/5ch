package handler

import (
	"ddd-sample/domain/model"
	"ddd-sample/usecase"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	GetAll() echo.HandlerFunc
	Get()    echo.HandlerFunc
	Post()   echo.HandlerFunc
	Patch()  echo.HandlerFunc
	Delete() echo.HandlerFunc
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
	ID        uint           `json:"id"`
	Name      model.UserName `json:"name"`
	Email     model.Email    `json:"email"`
	Password  model.Password `json:"password"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
}

type responseAllUser struct {
	Users *[]model.User `json:"users"`
}

func NewUserHandler(uu usecase.UserUsecase) UserHandler {
	return &userHandler{userUsecase: uu}
}

func RoutingUsers(e *echo.Echo, uh UserHandler) {
	e.GET("/users", uh.GetAll())
	e.GET("/users/:id", uh.Get())
	e.POST("/users", uh.Post())
	e.PATCH("/users/:id", uh.Patch())
	e.DELETE("/users/:id", uh.Delete())
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

func (uh *userHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSONPretty(http.StatusBadRequest, err.Error(), " ")
		}
		u, err := uh.userUsecase.Find(uint(id))
		if err != nil {
			return c.JSONPretty(http.StatusNotFound, err.Error(), " ")
		}
		res := allocateResponseUser(u)
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
		res := allocateResponseUser(u)
		return c.JSONPretty(http.StatusCreated, res, " ")
	}
}

func (uh *userHandler) Patch() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSONPretty(http.StatusBadRequest, err.Error(), " ")
		}
		req := new(requestUser)
		if err := c.Bind(req); err != nil {
			return c.JSONPretty(http.StatusBadRequest, err.Error(), " ")
		}
		if err := uh.userUsecase.Update(uint(id), req.Name, req.Email, req.Password); err != nil {
			return c.JSONPretty(http.StatusNotFound, err.Error(), " ")
		}
		return c.JSON(http.StatusNoContent, nil)
	}
}

func (uh *userHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSONPretty(http.StatusBadRequest, err.Error(), " ")
		}
		if err := uh.userUsecase.Delete(uint(id)); err != nil {
			return c.JSONPretty(http.StatusNotFound, err.Error(), " ")
		}
		return c.JSON(http.StatusNoContent, nil)
	}
}

func allocateResponseUser(u *model.User) *responseUser {
	res := &responseUser{
		ID:        u.GetID(),
		Name:      u.GetName(),
		Email:     u.GetEmail(),
		Password:  u.GetPassword(),
		CreatedAt: u.GetCreatedAt(),
		UpdatedAt: u.GetUpdatedAt(),
	}
	return res
}
