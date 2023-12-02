package routes

import (
	"api-gateway/controller"
	"api-gateway/middlewares"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, uc controller.UserController) {
	register := e.Group("/users/register")
	{
		register.POST("/user", uc.RegisterUser)
		register.POST("/driver", uc.RegisterDriver)
		register.POST("/admin", uc.RegisterAdmin)
	}
	e.POST("users/login", uc.Login)
	e.GET("users/logout", uc.Logout, middlewares.RequireAuth)
}