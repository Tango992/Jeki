package routes

import (
	"api-gateway/controller"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, uc controller.UserController) {
	register := e.Group("/users/register")
	{
		register.POST("/user", uc.RegisterUser)
		register.POST("/driver", uc.RegisterDriver)
		register.POST("/admin", uc.RegisterAdmin)
		register.POST("/login", uc.Login)
	}
}