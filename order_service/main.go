package main

import (
	"order-service/config"
	"order-service/controller"
	"order-service/repository"
)

func main() {
	db := config.ConnectDB().Database("jeki")
	orderCollection := db.Collection("orders")

	orderRepository := repository.NewOrderRepository(orderCollection)
	orderController := controller.NewOrderController(orderRepository)

	config.ListenAndServeGrpc(orderController)
}