package router

import (
	"api-gateway/controller"
	"api-gateway/middlewares"

	"github.com/labstack/echo/v4"
)

func Echo(e *echo.Echo, uc controller.UserController, mc controller.MerchantController, oc controller.OrderController) {
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

	restaurants := e.Group("/restaurants")
	{
		restaurants.GET("", mc.GetAllRestaurantsForCustomer)
		restaurants.GET("/:id", mc.GetRestaurantById)
		restaurants.GET("/:id", mc.GetMenuById)
	}

	merchant := e.Group("/merchant")
	merchant.Use(middlewares.RequireAuth)
	{
		restaurant := merchant.Group("/restaurant")
		{
			restaurant.GET("", mc.GetAllRestaurants)
			restaurant.POST("", mc.CreateRestaurant)
			restaurant.PUT("", mc.UpdateRestaurant)
		}
		menu := merchant.Group("/menu")
		{
			menu.GET("/:id", mc.GetOneMenuByAdminId)
			menu.POST("", mc.CreateMenu)
			menu.PUT("/:id", mc.UpdateMenu)
			menu.DELETE("/:id", mc.DeleteMenu)
		}
	}
	e.POST("maps", mc.TestMap)
}