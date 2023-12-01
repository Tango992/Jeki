package main

import (
	"log"
	"merchant-service/config"
	"merchant-service/controller"
	"merchant-service/middlewares"
	pb "merchant-service/pb/merchantpb"
	"merchant-service/repository"
	"net"
	"os"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(middlewares.NewInterceptorLogger()),
		),
	)

	merchantRepository := repository.NewMerchantRepository(db)
	merchantController := controller.NewUserController(merchantRepository)

	pb.RegisterMerchantServer(grpcServer, merchantController)

	port := os.Getenv("PORT")
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Server listening on port %s", port)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
