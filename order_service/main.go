package main

import "order-service/config"

func main() {
	db := config.ConnectDB().Database("jeki")
	orderCollection := db.Collection("orders")

	
}