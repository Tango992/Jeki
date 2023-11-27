package main

import (
	"api-gateway/helpers"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	e.Validator = &helpers.CustomValidator{NewValidator: validator.New()}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Logger.Fatal(e.Start(":8080"))
}