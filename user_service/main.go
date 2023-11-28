package main

import (
	"context"
	"fmt"
	"user-service/pb"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
) 

type Server struct {
	pb.UnimplementedUserServer
	db *gorm.DB
}

func NewServer() (*Server, error){
	dsn := ""
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	db.AutoMigrate(&pb.UserData{})

	return &Server{
		db: db,
	}, nil
}

func (s *Server) GetUserData (ctx context.Context, data *pb.EmailRequest) (*pb.UserData, error){
	var userData pb.UserData

	if err := s.db.Where("email = ?", data.Email).First(&userData).Error; err != nil {
		return nil, fmt.Errorf("failed to get user data: %v", err)
	}
	return &userData, nil
}

func (s *Server) Register (ctx context.Context, data *pb.RegisterRequest) (*pb.RegisterResponse, error){


}

func main() {
	
}