package main

import (
	"log"
	"merchant-service/config"
	"merchant-service/controller"
	"merchant-service/repository"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	merchantRepository := repository.NewMerchantRepository(db)
	merchantController := controller.NewMerchantController(merchantRepository)

	config.ListenAndServeGrpc(merchantController)
}
