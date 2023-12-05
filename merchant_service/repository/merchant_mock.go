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
	return []dto.MenuTmp{}, nil
}

func (m *MockMerchantRepository) FindRestaurantIdByAdminId(restaurantId uint32) (uint, error) {
	args := m.Mock.Called(restaurantId)
	return args.Get(0).(uint), args.Error(1)
}

func (m *MockMerchantRepository) FindRestaurantMetadataByMenuIds([]int) (*pb.RestaurantMetadata, error) {
	return &pb.RestaurantMetadata{}, nil
}

func (m *MockMerchantRepository) FindAdminIdByMenuId(adminID uint32) (uint32, error) {
	return adminID, nil
}

func (m *MockMerchantRepository) FindMenuByRestaurantId(restaurantId uint32) ([]*pb.Menu, error) {
	args := m.Mock.Called(restaurantId)
	return args.Get(0).([]*pb.Menu), args.Error(1)
}

func (m *MockMerchantRepository) FindMenuById(menuId uint32) (*pb.Menu, error) {
	return &pb.Menu{}, nil
}

func (m *MockMerchantRepository) FindMenusByAdminId(adminId uint32) ([]*pb.MenuCompact, error) {
	return []*pb.MenuCompact{}, nil
}

func (m *MockMerchantRepository) FindOneMenuByAdminId(menuID uint32, adminID uint32) (*pb.MenuCompact, error) {
	return &pb.MenuCompact{}, nil
}

func (m *MockMerchantRepository) FindRestaurantByAdminId(adminID uint32) (*pb.RestaurantData, error) {
	return &pb.RestaurantData{}, nil
}

func (m *MockMerchantRepository) UpdateMenu(*pb.UpdateMenuData) error {
	return nil
}

func (m *MockMerchantRepository) UpdateRestaurant(restaurantId uint, data *pb.UpdateRestaurantData) error {
	return nil
}

func (m *MockMerchantRepository) DeleteMenu(uint, uint) error {
	return nil
}

func (m *MockMerchantRepository) CreateMenu(menu *model.Menu) error {
	args := m.Mock.Called(menu) 
	return args.Error(0)
}

func (m *MockMerchantRepository) CreateRestaurant(*model.Restaurant) error {
	return nil
}
