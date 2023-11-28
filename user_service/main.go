package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"user-service/pb"

	"google.golang.org/grpc"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	pb.UnimplementedUserServer
	db *gorm.DB
}

func NewServer() (*Server, error) {
	dsn := "user=postgres dbname=deploy host=localhost port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	db.AutoMigrate(&pb.UserData{})

	return &Server{
		db: db,
	}, nil
}

func (s *Server) GetUserData(ctx context.Context, data *pb.EmailRequest) (*pb.UserData, error) {
	var userData pb.UserData
	result := s.db.First(&userData, "email = ?", data.Email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &userData, nil
}

func (s *Server) Register(ctx context.Context, data *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	newUser := pb.UserData{
		Email:    data.Email,
		Password: data.Password,
	}

	result := s.db.Create(&newUser)
	if result.Error != nil {
		return nil, result.Error
	}

	userID := uint32(1)
	createdAt := "2023-11-28T12:00:00"

	response := &pb.RegisterResponse{
		UserId:    userID,
		CreatedAt: createdAt,
	}

	return response, nil

}

func main() {
	server, err := NewServer()
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register the UserServer with the gRPC server
	pb.RegisterUserServer(grpcServer, server)

	// Set up a listener on a TCP port
	port := 50051
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Server listening on port %d", port)

	// Start the gRPC server
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
