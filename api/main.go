package main

import (
	"ddd-sample/config"
	"ddd-sample/infra"
	"ddd-sample/interface/handler"
	"ddd-sample/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	userRepository := infra.NewUserRepository(config.InitDB())
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	e := echo.New()
	handler.RoutingUsers(e, userHandler)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		ContentTypeNosniff: "application/json",
	}))

	e.Start(":8080")
}
