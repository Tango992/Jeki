package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"
	"user-service/models"
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

	if err := db.AutoMigrate(&models.Role{}, &models.User{}, &models.Verification{}); err != nil {
		log.Fatal(err)
	}

	return &Server{
		db: db,
	}, nil
}

/*
	"email": "john@mail.com",
	"password": "secret"
*/

func convertUserToUserData(user models.User) *pb.UserData {
	return &pb.UserData{
		Id:        uint32(user.ID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
		BirthDate: user.BirthDate,
		// Role:      user.Role,
	}
}

func (s *Server) GetUserData(ctx context.Context, data *pb.EmailRequest) (*pb.UserData, error) {
	var user models.User
	result := s.db.First(&user, "email = ?", data.Email)
	if result.Error != nil {
		return nil, result.Error
	}

	userData := convertUserToUserData(user)

	return userData, nil
}

func (s *Server) Register(ctx context.Context, data *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	newUser := models.User{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Password:  data.Password,
		BirthDate: data.BirthDate,
		RoleID:    data.RoleId,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	result := s.db.Create(&newUser)
	if result.Error != nil {
		return nil, result.Error
	}

	response := &pb.RegisterResponse{
		UserId:    uint32(newUser.ID),
		CreatedAt: newUser.CreatedAt,
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
