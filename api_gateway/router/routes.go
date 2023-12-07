package router

import (
	"api-gateway/controller"
	_ "api-gateway/docs"
	"api-gateway/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
)

func Echo(e *echo.Echo, uc controller.UserController, mc controller.MerchantController, oc controller.OrderController) {
	e.GET("/user/verified", func(c echo.Context) error {
		return c.Render(http.StatusOK, "verification.html", nil)
	})
	
	e.GET("", func(c echo.Context) error {
		return c.Redirect(http.StatusTemporaryRedirect, "/swagger/index.html")
	})
	e.GET("/swagger/*", echoSwagger.WrapHandler)

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

	restaurants := e.Group("/restaurant")
	{
		restaurants.GET("", mc.GetAllRestaurantsForCustomer)
		restaurants.GET("/:id", mc.GetRestaurantById)
	}
	e.GET("menu/:id", mc.GetMenuById)

	merchant := e.Group("/merchant")
	merchant.Use(middlewares.RequireAuth)
	{
		restaurant := merchant.Group("/restaurant")
		{
			restaurant.GET("", mc.GetRestaurantByAdminId)
			restaurant.POST("", mc.CreateRestaurant)
			restaurant.PUT("", mc.UpdateRestaurant)
		}
		menu := merchant.Group("/menu")
		{
			menu.GET("", mc.GetMenuByAdminId)
			menu.GET("/:id", mc.GetOneMenuByAdminId)
			menu.POST("", mc.CreateMenu)
			menu.PUT("/:id", mc.UpdateMenu)
			menu.DELETE("/:id", mc.DeleteMenu)
		}
	}

	order := e.Group("")
	order.Use(middlewares.RequireAuth)
	{	
		users := order.Group("/users")
		{
			users.POST("/orders", oc.UserCreateOrder)
			users.GET("/orders", oc.UsersGetAllOrders)
			users.GET("/ongoing", oc.UsersGetOngoingOrder)
			users.GET("/orders/:id", oc.UsersGetOrderById)
		}

		merchant := order.Group("/merchant")
		{
			merchant.GET("/orders", oc.MerchantGetAllOrders)
			merchant.GET("/ongoing", oc.MerchantGetOngoingOrder)
			merchant.GET("/orders/:id", oc.MerchantGetOrderById)
			merchant.PUT("/orders/:id", oc.MerchantUpdateOrder)
		}

		driver := order.Group("/driver")
		{
			driver.GET("/orders", oc.DriverGetAllOrders)
			driver.GET("/ongoing", oc.DriverGetCurrentOrder)
			driver.GET("/orders/:id", oc.DriverGetOrderById)
			driver.PUT("/orders/:id", oc.DriverUpdateOrder)
		}
	}
	
	payment := e.Group("/payment")
	{
		payment.POST("/xendit/update", oc.PaymentUpdate)
	}
}