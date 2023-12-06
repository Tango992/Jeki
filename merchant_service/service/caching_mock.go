package service

import (
	"merchant-service/pb/merchantpb"

	"github.com/stretchr/testify/mock"
)

type MockCachingService struct {
	Mock mock.Mock
}

func NewMockCachingService() MockCachingService {
	return MockCachingService{}
}

func (m *MockCachingService) SetRestaurantDetailed(restaurantId uint, data any) error {
	args := m.Mock.Called(restaurantId, data)
	return args.Error(0)
}

func (m *MockCachingService) GetRestaurantDetailed(restaurantId uint) (*merchantpb.RestaurantDetailed, error){
	args := m.Mock.Called(restaurantId)
	return args.Get(0).(*merchantpb.RestaurantDetailed), args.Error(1)
}