package controller_test

import (
	"context"
	"testing"
	"time"
	"user-service/controller"
	"user-service/dto"
	"user-service/helpers"
	"user-service/models"
	"user-service/pb/userpb"
	"user-service/repository"
	"user-service/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	OnlineStatus  = "available"
	OngoingStatus = "ongoing"
	OfflineStatus = "offline"
)

var (
	mockRepository = repository.NewMockUserRepository()
	mockMessageBroker = service.NewMockMessageBroker()
	userController = controller.NewUserController(&mockRepository, &mockMessageBroker)
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestGetUserData(t *testing.T) {
	user := dto.UserJoinedData{
		ID: 1,
		FirstName: "John",
		LastName: "Doe",
		Email: "example@example.com",
		Password: "secret",
		BirthDate: "2001-01-01",
		Role: "user",
		CreatedAt: "2023-12-05",
		Verified: true,
	}

	pbRequestData := &userpb.EmailRequest{
		Email: "example@example.com",
	}

	mockRepository.Mock.On("GetUserData", pbRequestData.Email).Return(user, error(nil))
	pbResponse, err := userController.GetUserData(context.Background(), pbRequestData)
	responseDummy := helpers.ConvertUserToUserData(user)
	
	assert.Nil(t, err)
	assert.NotEmpty(t, pbResponse)
	assert.Equal(t, responseDummy, pbResponse)
}

func TestRegister(t *testing.T) {
	var dummyUserID uint = 1
	dummyToken := helpers.GenerateVerificationToken()

	pbRequestData := &userpb.RegisterRequest{
		FirstName: "John",
		LastName: "Doe",
		Email: "example@example.com",
		Password: "secret",
		BirthDate: "2001-01-01",
		RoleId: 1,
	}

	newUser := models.User{
		FirstName: pbRequestData.FirstName,
		LastName:  pbRequestData.LastName,
		Email:     pbRequestData.Email,
		Password:  pbRequestData.Password,
		BirthDate: pbRequestData.BirthDate,
		RoleID:    pbRequestData.RoleId,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	verificationData := models.Verification{
		UserID: newUser.ID,
		Token:  dummyToken,
	}

	pbResponse := &userpb.RegisterResponse{
		UserId: uint32(dummyUserID),
		CreatedAt: newUser.CreatedAt,
	}

	mockRepository.Mock.On("CreateUser", &newUser).Return(error(nil)).Run(func(args mock.Arguments) {
		userData := args.Get(0).(*models.User)
		userData.ID = dummyUserID
	})
	
	mockRepository.Mock.On("AddToken", &verificationData).Return(error(nil))

	mockMessageBroker.Mock.On("PublishMessage", mock.Anything).Return(error(nil))
	
	pbResponse, err := userController.Register(context.Background(), pbRequestData)
	assert.Nil(t, err)
	assert.NotEmpty(t, pbResponse)
	assert.Equal(t, pbResponse.UserId, uint32(dummyUserID))
}

func TestGetAvailableDriver(t *testing.T) {
	var mockDriverID uint32 = 3
	driverData := dto.DriverData{
		ID: int(mockDriverID),
		Name: "John Doe",
	}

	mockRepository.Mock.On("GetAvailableDriver").Return(driverData, error(nil))
	
	pbResponse, err := userController.GetAvailableDriver(context.Background(), &emptypb.Empty{})
	assert.Nil(t, err)
	assert.NotEmpty(t, pbResponse)
	assert.Equal(t, mockDriverID, pbResponse.Id)
}

func TestSetDriverStatusOnline(t *testing.T) {
	var mockDriverID uint32 = 3
	pbRequest := &userpb.DriverId{
		Id: mockDriverID,
	}

	mockRepository.Mock.On("SetDriverStatusOnline", uint(mockDriverID)).Return(error(nil))
	
	pbResponse, err := userController.SetDriverStatusOnline(context.Background(), pbRequest)
	assert.Nil(t, err)
	assert.Empty(t, pbResponse)
}

func TestSetDriverStatusOngoing(t *testing.T) {
	var mockDriverID uint32 = 3
	pbRequest := &userpb.DriverId{
		Id: mockDriverID,
	}

	mockRepository.Mock.On("SetDriverStatusOngoing", uint(mockDriverID)).Return(error(nil))

	pbResponse, err := userController.SetDriverStatusOngoing(context.Background(), pbRequest)
	assert.Nil(t, err)
	assert.Empty(t, pbResponse)
}

func TestSetDriverStatusOffline(t *testing.T) {
	var mockDriverID uint32 = 3
	pbRequest := &userpb.DriverId{
		Id: mockDriverID,
	}

	mockRepository.Mock.On("SetDriverStatusOffline", uint(mockDriverID)).Return(error(nil))

	pbResponse, err := userController.SetDriverStatusOffline(context.Background(), pbRequest)
	assert.Nil(t, err)
	assert.Empty(t, pbResponse)
}

func TestVerifyNewUser(t *testing.T) {
	var mockUserID uint32 = 1
	mockToken := "abcdefg"
	
	pbRequest := &userpb.UserCredential{
		Id: mockUserID,
		Token: mockToken,
	}

	mockRepository.Mock.On("VerifyNewUser", pbRequest.Id, pbRequest.Token).Return(error(nil))

	pbResponse, err := userController.VerifyNewUser(context.Background(), pbRequest)
	assert.Nil(t, err)
	assert.Empty(t, pbResponse)
}

func TestCreateDriverData(t *testing.T) {
	var mockUserID uint32 = 1

	pbRequest := &userpb.DriverId{
		Id: mockUserID,
	}

	mockRepository.Mock.On("CreateDriverData", mockUserID).Return(error(nil))
	
	pbResponse, err := userController.CreateDriverData(context.Background(), pbRequest)
	assert.Nil(t, err)
	assert.Empty(t, pbResponse)
}