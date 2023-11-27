package main

import (
	"api-gateway/config"
	"api-gateway/controller"
	"api-gateway/helpers"
	"api-gateway/routes"

	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	conn, userClient := config.InitGrpc()
	defer conn.Close()

	userController := controller.NewUserController(userClient)
	
	e.Validator = &helpers.CustomValidator{NewValidator: validator.New()}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.Routes(e, userController)

	e.Logger.Fatal(e.Start(":8080"))
}