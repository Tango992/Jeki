package main

import (
	"log"
	"user-service/config"
	"user-service/controller"
	"user-service/repository"
	"user-service/service"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	conn, mbChan := config.InitMessageBroker()
	defer conn.Close()
	
	messageBrokerService := service.NewMessageBroker(mbChan)
	userRepository := repository.NewUserRepository(db)
	userController := controller.NewUserController(userRepository, messageBrokerService)
	
	config.ListenAndServeGrpc(userController)
}
