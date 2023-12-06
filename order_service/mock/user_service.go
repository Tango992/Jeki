package mock

import (
	"context"
	"order-service/pb/userpb"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MockUserService struct {
	Mock mock.Mock
}

func NewMockUserService() MockUserService {
	return MockUserService{}
}

func (m *MockUserService) Register(ctx context.Context, in *userpb.RegisterRequest, opts ...grpc.CallOption) (*userpb.RegisterResponse, error) {
	return nil,  nil
}

func (m *MockUserService) GetUserData(ctx context.Context, in *userpb.EmailRequest, opts ...grpc.CallOption) (*userpb.UserData, error) {
	return nil, nil
}

func (m *MockUserService) GetAvailableDriver(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*userpb.DriverData, error) {
	args := m.Mock.Called(ctx, in)
	return args.Get(0).(*userpb.DriverData), args.Error(1)
}

func (m *MockUserService) CreateDriverData(ctx context.Context, in *userpb.DriverId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *MockUserService) SetDriverStatusOnline(ctx context.Context, in *userpb.DriverId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	args := m.Mock.Called(ctx, in)
	return args.Get(0).(*emptypb.Empty), args.Error(1)
}

func (m *MockUserService) SetDriverStatusOngoing(ctx context.Context, in *userpb.DriverId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	args := m.Mock.Called(ctx, in)
	return args.Get(0).(*emptypb.Empty), args.Error(1)
}

func (m *MockUserService) SetDriverStatusOffline(ctx context.Context, in *userpb.DriverId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	args := m.Mock.Called(ctx, in)
	return args.Get(0).(*emptypb.Empty), args.Error(1)
}

func (m *MockUserService) VerifyNewUser(ctx context.Context, in *userpb.UserCredential, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}