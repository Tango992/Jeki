package controller

import (
	"context"
	"time"
	"user-service/dto"
	"user-service/helpers"
	"user-service/models"
	"user-service/pb"
	"user-service/repository"
)

type Server struct {
	pb.UnimplementedUserServer
	Repository repository.User
}

func NewUserController(r repository.User) Server {
	return Server{
		Repository: r,
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

func (s Server) GetUserData(ctx context.Context, data *pb.EmailRequest) (*pb.UserData, error) {
	user, err := s.Repository.GetUserData(data.Email)
	if err != nil {
		return nil, err
	}

	userData := convertUserToUserData(user)

	return userData, nil
}

func (s Server) Register(ctx context.Context, data *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	newUser := models.User{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Password:  data.Password,
		BirthDate: data.BirthDate,
		RoleID:    data.RoleId,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	if err := s.Repository.CreateUser(&newUser); err != nil {
		return nil, err
	}
	
	verificationData := models.Verification{
		UserID: newUser.ID,
		Token: helpers.GenerateVerificationToken(),
	}
	
	if err := s.Repository.AddToken(&verificationData); err != nil {
		return nil, err
	}

	response := &pb.RegisterResponse{
		UserId:    uint32(newUser.ID),
		CreatedAt: newUser.CreatedAt,
	}

	return response, nil
}
