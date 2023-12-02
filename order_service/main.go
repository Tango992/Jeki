package main

import (
	"order-service/config"
	"order-service/controller"
	"order-service/repository"
	"order-service/service"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db := config.ConnectDB().Database("jeki")
	orderCollection := db.Collection("orders")

	conn, userService := config.InitUserService()
	defer conn.Close()
	
	paymentService := service.NewPaymentService(os.Getenv("XENDIT_API_KEY"))
	orderRepository := repository.NewOrderRepository(orderCollection)
	orderController := controller.NewOrderController(orderRepository, paymentService, userService)

	config.ListenAndServeGrpc(orderController)
}