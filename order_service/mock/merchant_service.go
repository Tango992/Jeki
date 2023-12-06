package mock

import (
	"context"
	"order-service/pb/merchantpb"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MockMerchantService struct {
	Mock mock.Mock
}

func NewMockMerchantService() MockMerchantService {
	return MockMerchantService{}
}

func (m *MockMerchantService) FindAllRestaurants(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*merchantpb.RestaurantCompactRepeated, error) {
	return nil, nil
}

func (m *MockMerchantService) FindRestaurantById(ctx context.Context, in *merchantpb.IdRestaurant, opts ...grpc.CallOption) (*merchantpb.RestaurantDetailed, error) {
	return nil, nil
}

func (m *MockMerchantService) FindMenuById(ctx context.Context, in *merchantpb.MenuId, opts ...grpc.CallOption) (*merchantpb.Menu, error) {
	return nil, nil
}

func (m *MockMerchantService) CalculateOrder(ctx context.Context, in *merchantpb.RequestMenuDetails, opts ...grpc.CallOption) (*merchantpb.CalculateOrderResponse, error) {
	args := m.Mock.Called(ctx, in)
	return args.Get(0).(*merchantpb.CalculateOrderResponse), args.Error(1)
}

func (m *MockMerchantService) CreateRestaurant(ctx context.Context, in *merchantpb.NewRestaurantData, opts ...grpc.CallOption) (*merchantpb.IdRestaurant, error) {
	return nil, nil
}

func (m *MockMerchantService) UpdateRestaurant(ctx context.Context, in *merchantpb.UpdateRestaurantData, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *MockMerchantService) CreateMenu(ctx context.Context, in *merchantpb.NewMenuData, opts ...grpc.CallOption) (*merchantpb.MenuId, error) {
	return nil, nil
}

func (m *MockMerchantService) UpdateMenu(ctx context.Context, in *merchantpb.UpdateMenuData, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *MockMerchantService) DeleteMenu(ctx context.Context, in *merchantpb.AdminIdMenuId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *MockMerchantService) FindRestaurantByAdminId(ctx context.Context, in *merchantpb.AdminId, opts ...grpc.CallOption) (*merchantpb.RestaurantData, error) {
	return nil, nil
}

func (m *MockMerchantService) FindMenusByAdminId(ctx context.Context, in *merchantpb.AdminId, opts ...grpc.CallOption) (*merchantpb.MenuCompactRepeated, error) {
	return nil, nil
}

func (m *MockMerchantService) FindOneMenuByAdminId(ctx context.Context, in *merchantpb.AdminIdMenuId, opts ...grpc.CallOption) (*merchantpb.MenuCompact, error) {
	return nil, nil
}