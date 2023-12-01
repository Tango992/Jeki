package controller

import (
	"context"
	"encoding/json"
	"time"
	"user-service/dto"
	"user-service/helpers"
	"user-service/models"
	pb "user-service/pb/userpb"
	"user-service/repository"
	"user-service/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	pb.UnimplementedUserServer
	Repository repository.User
	Mb         service.MessageBroker
}

func NewUserController(r repository.User, mb service.MessageBroker) Server {
	return Server{
		Repository: r,
		Mb:         mb,
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
		Token:  helpers.GenerateVerificationToken(),
	}

	if err := s.Repository.AddToken(&verificationData); err != nil {
		return nil, err
	}

	dataJsonRequest := dto.UserMessageBroker{
		ID:    newUser.ID,
		Name:  newUser.FirstName,
		Email: newUser.Email,
		Token: verificationData.Token,
	}

	dataJson, err := json.Marshal(dataJsonRequest)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err := s.Mb.PublishMessage(dataJson); err != nil {
		return nil, err
	}

	response := &pb.RegisterResponse{
		UserId:    uint32(newUser.ID),
		CreatedAt: newUser.CreatedAt,
	}

	return response, nil
}

func (s Server) GetAvailableDriver(ctx context.Context, data *emptypb.Empty) (*pb.DriverData, error) {
	driver, err := s.Repository.GetAvailableDriver()
	if err != nil {
		return nil, err
	}

	driverData := &pb.DriverData{
		Id:   uint32(driver.ID),
		Name: driver.Name,
	}

	return driverData, nil
}
