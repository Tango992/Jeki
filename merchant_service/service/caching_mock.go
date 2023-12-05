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
	return nil
}

func (m *MockCachingService) GetRestaurantDetailed(restaurantId uint) (*merchantpb.RestaurantDetailed, error){
	return nil, nil
}