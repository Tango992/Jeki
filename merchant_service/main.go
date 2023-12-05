package main

import (
	"log"
	"merchant-service/config"
	"merchant-service/controller"
	"merchant-service/repository"
	"merchant-service/service"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	redisClient := config.InitRedisClient()
	cachingService := service.NewCachingService(redisClient)

	merchantRepository := repository.NewMerchantRepository(db)
	merchantController := controller.NewMerchantController(merchantRepository, cachingService)

	config.ListenAndServeGrpc(merchantController)
}
