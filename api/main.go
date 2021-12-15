package main

import (
	"bbs/config"
	"bbs/infrastructure"
	"bbs/interface/handler"
	"bbs/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	userRepository := infrastructure.NewUserRepository(config.NewDB())
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	e := echo.New()
	handler.RoutingUsers(e, userHandler)
	e.Start(":8080")
}
