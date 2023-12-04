package routes

import (
	"api-gateway/controller"
	"api-gateway/middlewares"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, uc controller.UserController, mc controller.MerchantController, oc controller.OrderController) {
	users := e.Group("/users")
	{
		register := users.Group("/register")
		{
			register.POST("/user", uc.RegisterUser)
			register.POST("/driver", uc.RegisterDriver)
			register.POST("/admin", uc.RegisterAdmin)
		}
		
		users.POST("/login", uc.Login)
		users.GET("/verify/:userid/:token", uc.VerifyUser)
		users.GET("/logout", uc.Logout, middlewares.RequireAuth)
	}
}