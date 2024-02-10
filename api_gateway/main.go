package main

import (
	"api-gateway/config"
	"api-gateway/controller"
	"api-gateway/helpers"
	"api-gateway/router"
	"api-gateway/service"
	"html/template"
	"io"
	"os"

	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
    templates *template.Template
}

// @title Jeki
// @version 1.0
// @description Food delivery app built with microservices that integrates customer, driver, and restaurant.

// @contact.name Contact the developer
// @contact.email daniel.rahmanto@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host jeki-egmflbdzpa-et.a.run.app
// @BasePath /
func main() {
	e := echo.New()
	e.Validator = &helpers.CustomValidator{NewValidator: validator.New()}
	e.Renderer = &Template{
		templates: template.Must(template.ParseGlob("./template/*.html")),
	}
	e.Static("/template", "template/")
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

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}