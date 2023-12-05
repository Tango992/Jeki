package repository

import (
	"user-service/dto"
	"user-service/models"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	Mock mock.Mock
}

func NewMockUserRepository() MockUserRepository {
	return MockUserRepository{}
}

func (m *MockUserRepository) GetUserData(email string) (dto.UserJoinedData, error) {
	args := m.Mock.Called(email)
	return args.Get(0).(dto.UserJoinedData), args.Error(1)
}

func (m *MockUserRepository) CreateUser(data *models.User) error {
	args := m.Mock.Called(data)
	return args.Error(0)
}

func (m *MockUserRepository) CreateDriverData(userID uint32) error {
	args := m.Mock.Called(userID)
	return args.Error(0)
}

func (m *MockUserRepository) AddToken(data *models.Verification) error {
	args := m.Mock.Called(mock.Anything)
	return args.Error(0)
}

func (m *MockUserRepository) GetAvailableDriver() (dto.DriverData, error) {
	args := m.Mock.Called()
	return args.Get(0).(dto.DriverData), args.Error(1)
}

func (m *MockUserRepository) SetDriverStatusOnline(driverID uint) error {
	args := m.Mock.Called(driverID)
	return args.Error(0)
}

func (m *MockUserRepository) SetDriverStatusOngoing(driverID uint) error {
	args := m.Mock.Called(driverID)
	return args.Error(0)
}

func (m *MockUserRepository) SetDriverStatusOffline(driverID uint) error {
	args := m.Mock.Called(driverID)
	return args.Error(0)
}

func (m *MockUserRepository) VerifyNewUser(id uint32, token string) error {
	args := m.Mock.Called(id, token)
	return args.Error(0)
}