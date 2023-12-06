package repository

import (
	"merchant-service/dto"
	"merchant-service/model"
	pb "merchant-service/pb/merchantpb"

	"github.com/stretchr/testify/mock"
)

type MockMerchantRepository struct {
	Mock mock.Mock
}

func NewMockMerchantRepository() MockMerchantRepository {
	return MockMerchantRepository{}
}

func (m *MockMerchantRepository) FindAllRestaurants() ([]model.Restaurant, error) {
	args := m.Mock.Called()
	return args.Get(0).([]model.Restaurant), args.Error(1)
}

func (m *MockMerchantRepository) FindRestaurantByID(id uint32) (dto.RestaurantDataCompact, error) {
	args := m.Mock.Called(id)
	return args.Get(0).(dto.RestaurantDataCompact), args.Error(1)
}

func (m *MockMerchantRepository) FindMultipleMenuDetails(menuIds []int) ([]dto.MenuTmp, error) {
	args := m.Mock.Called(menuIds)
	return args.Get(0).([]dto.MenuTmp), args.Error(1)
}

func (m *MockMerchantRepository) FindRestaurantIdByAdminId(restaurantId uint32) (uint, error) {
	args := m.Mock.Called(restaurantId)
	return args.Get(0).(uint), args.Error(1)
}

func (m *MockMerchantRepository) FindRestaurantMetadataByMenuIds(menuIDs []int) (*pb.RestaurantMetadata, error) {
	args := m.Mock.Called(menuIDs)
	return args.Get(0).(*pb.RestaurantMetadata), args.Error(1)
}

func (m *MockMerchantRepository) FindAdminIdByMenuId(adminID uint32) (uint32, error) {
	args := m.Mock.Called(adminID)
	return args.Get(0).(uint32), args.Error(1)
}

func (m *MockMerchantRepository) FindMenuByRestaurantId(restaurantId uint32) ([]*pb.Menu, error) {
	args := m.Mock.Called(restaurantId)
	return args.Get(0).([]*pb.Menu), args.Error(1)
}

func (m *MockMerchantRepository) FindMenuById(menuId uint32) (*pb.Menu, error) {
	args := m.Mock.Called(menuId)
	return args.Get(0).(*pb.Menu), args.Error(1)
}

func (m *MockMerchantRepository) FindMenusByAdminId(adminId uint32) ([]*pb.MenuCompact, error) {
	args := m.Mock.Called(adminId)
	return args.Get(0).([]*pb.MenuCompact), args.Error(1)
}

func (m *MockMerchantRepository) FindOneMenuByAdminId(menuID uint32, adminID uint32) (*pb.MenuCompact, error) {
	args := m.Mock.Called(menuID, adminID)
	return args.Get(0).(*pb.MenuCompact), args.Error(1)
}

func (m *MockMerchantRepository) FindRestaurantByAdminId(adminID uint32) (*pb.RestaurantData, error) {
	args := m.Mock.Called(adminID)
	return args.Get(0).(*pb.RestaurantData), args.Error(1)
}

func (m *MockMerchantRepository) UpdateMenu(data *pb.UpdateMenuData) error {
	args := m.Mock.Called(data)
	return args.Error(0)
}

func (m *MockMerchantRepository) UpdateRestaurant(restaurantId uint, data *pb.UpdateRestaurantData) error {
	args := m.Mock.Called(restaurantId, data)
	return args.Error(0)
}

func (m *MockMerchantRepository) DeleteMenu(restaurantID, menuID uint) error {
	args := m.Mock.Called(restaurantID, menuID)
	return args.Error(0)
}

func (m *MockMerchantRepository) CreateMenu(menu *model.Menu) error {
	args := m.Mock.Called(menu) 
	return args.Error(0)
}

func (m *MockMerchantRepository) CreateRestaurant(data *model.Restaurant) error {
	args := m.Mock.Called(data)
	return args.Error(0)
}
