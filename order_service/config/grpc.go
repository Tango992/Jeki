package config

import (
	"log"
	"net"
	"os"
	"order-service/controller"
	"order-service/middlewares"
	"order-service/pb"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"
)

func ListenAndServeGrpc(controller controller.OrderController) {
	port := os.Getenv("PORT")
	
	lis, err := net.Listen("tcp", ":" + port)
	if err != nil {
		log.Fatal(err)
	}
	
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(middlewares.NewInterceptorLogger()),
		),
	)

	pb.RegisterOrderServiceServer(grpcServer, controller)

	log.Println("\033[36mGRPC server is running on port:", port, "\033[0m")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to server gRPC:", err)
	}
}