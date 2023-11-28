package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"time"
	"user-service/dto"
	"user-service/config"
	"user-service/models"
	"user-service/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gorm.io/gorm"
)

type Server struct {
	pb.UnimplementedUserServer
	db *gorm.DB
}


func NewServer(db *gorm.DB) *Server {
	if err := db.AutoMigrate(&models.Role{}, &models.User{}, &models.Verification{}); err != nil {
		log.Fatal(err)
	}

	return &Server{
		db: db,
	}
}


func convertUserToUserData(user dto.UserJoinedData) *pb.UserData {
	return &pb.UserData{
		Id:        uint32(user.ID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
		BirthDate: user.BirthDate,
		Role:      user.Role,
	}
}

func (s *Server) GetUserData(ctx context.Context, data *pb.EmailRequest) (*pb.UserData, error) {
	var user dto.UserJoinedData

	result := s.db.Table("users u").
		Select("u.id, u.first_name, u.last_name, u.email, u.password, u.birth_date, u.created_at, r.name AS role").
		Where("u.email = ?", data.Email).
		Joins("JOIN roles r on u.role_id = r.id").
		Take(&user)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal,err.Error())
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
	if err := result.Error; err != nil {
		if err.Error() == `ERROR: duplicate key value violates unique constraint "users_email_key" (SQLSTATE 23505)` {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}
		return nil, status.Error(codes.Internal,err.Error())
	}

	response := &pb.RegisterResponse{
		UserId:    uint32(newUser.ID),
		CreatedAt: newUser.CreatedAt,
	}

	return response, nil
}

func main() {
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get underlying DB: %v", err)
	}

	defer sqlDB.Close()

	server := NewServer(db)

	grpcServer := grpc.NewServer()

	pb.RegisterUserServer(grpcServer, server)

	port := 50051
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Server listening on port %d", port)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
