package main

import (
	"api-gateway/config"
	"api-gateway/controller"
	"api-gateway/helpers"
	"api-gateway/router"
	"api-gateway/service"
	"os"

	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Validator = &helpers.CustomValidator{NewValidator: validator.New()}
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	conn, userClient := config.InitUserServiceClient()
	defer conn.Close()

	conn, merchantClient := config.InitMerchantServiceClient()
	defer conn.Close()

	conn, orderClient := config.InitOrderServiceClient()
	defer conn.Close()

	mapsService := service.NewMapsService()
	userController := controller.NewUserController(userClient)
	merchantController := controller.NewMerchantController(merchantClient, mapsService)
	orderController := controller.NewOrderController(orderClient, mapsService)
	
	router.Echo(e, userController, merchantController, orderController)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}