package main

import (
	"log"
	"net"
	"os"
	"user-service/config"
	"user-service/controller"
	"user-service/middlewares"
	"user-service/pb"
	"user-service/repository"

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
	
	userRepository := repository.NewUserRepository(db)
	userController := controller.NewUserController(userRepository)
	
	pb.RegisterUserServer(grpcServer, userController)

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
