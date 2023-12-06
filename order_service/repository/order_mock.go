package repository

import (
	"context"
	"order-service/model"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockOrderRepository struct {
	Mock mock.Mock
}

func NewMockOrderRepository() MockOrderRepository {
	return MockOrderRepository{}
}

func (m *MockOrderRepository) Create(ctx context.Context, data *model.Order) error {
	args := m.Mock.Called(ctx, data)
	return args.Error(0)
}

func (m *MockOrderRepository) FindById(ctx context.Context, id string) (model.Order, error) {
	args := m.Mock.Called(ctx, id)
	return args.Get(0).(model.Order), args.Error(1)
}

func (m *MockOrderRepository) FindWithFilter(ctx context.Context, filter bson.D) ([]model.Order, error) {
	return []model.Order{}, nil
}

func (m *MockOrderRepository) UpdateWithFilter(ctx context.Context, field bson.D, data bson.M) error {
	return nil
}

func (m *MockOrderRepository) CompleteOrderStatus(ctx context.Context, orderId primitive.ObjectID) error {
	args := m.Mock.Called(ctx, orderId)
	return args.Error(0)
}

func (m *MockOrderRepository) CancelOrderStatus(ctx context.Context, orderId primitive.ObjectID) error {
	args := m.Mock.Called(ctx, orderId)
	return args.Error(0)
}

func (m *MockOrderRepository) FindRestaurantAllOrders(ctx context.Context, adminID uint) ([]model.Order, error) {
	args := m.Mock.Called(ctx, adminID)
	return args.Get(0).([]model.Order), args.Error(1)
}

func (m *MockOrderRepository) FindRestaurantCurrentOrders(ctx context.Context, adminID uint) ([]model.Order, error) {
	args := m.Mock.Called(ctx, adminID)
	return args.Get(0).([]model.Order), args.Error(1)
}

func (m *MockOrderRepository) FindUserAllOrders(ctx context.Context, userID uint) ([]model.Order, error) {
	args := m.Mock.Called(ctx, userID)
	return args.Get(0).([]model.Order), args.Error(1)
}

func (m *MockOrderRepository) FindUserCurrentOrders(ctx context.Context, userID uint) ([]model.Order, error) {
	args := m.Mock.Called(ctx, userID)
	return args.Get(0).([]model.Order), args.Error(1)
}

func (m *MockOrderRepository) FindDriverAllOrders(ctx context.Context, driverID uint) ([]model.Order, error) {
	args := m.Mock.Called(ctx, driverID)
	return args.Get(0).([]model.Order), args.Error(1)
}

func (m *MockOrderRepository) FindDriverCurrentOrder(ctx context.Context, driverID uint) (model.Order, error) {
	args := m.Mock.Called(ctx, driverID)
	return args.Get(0).(model.Order), args.Error(1)
}

func (m *MockOrderRepository) UpdateRestaurantStatus(ctx context.Context, orderId primitive.ObjectID, userId uint32, status string) error {
	args := m.Mock.Called(ctx, orderId, userId, status)
	return args.Error(0)
}

func (m *MockOrderRepository) UpdateDriverStatus(ctx context.Context, orderId primitive.ObjectID, userId uint32, status string) error {
	return nil
}

func (m *MockOrderRepository) UpdatePaymentStatus(ctx context.Context, orderId primitive.ObjectID, status, method, completedAt string) error {
	args := m.Mock.Called(ctx, orderId, status, method, completedAt)
	return args.Error(0)
}