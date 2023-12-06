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
	return nil
}

func (m *MockOrderRepository) FindById(ctx context.Context, id string) (model.Order, error) {
	return model.Order{}, nil
}

func (m *MockOrderRepository) FindWithFilter(ctx context.Context, filter bson.D) ([]model.Order, error) {
	return []model.Order{}, nil
}

func (m *MockOrderRepository) UpdateWithFilter(ctx context.Context, field bson.D, data bson.M) error {
	return nil
}

func (m *MockOrderRepository) CompleteOrderStatus(ctx context.Context, orderId primitive.ObjectID) error {
	return nil
}

func (m *MockOrderRepository) CancelOrderStatus(ctx context.Context, orderId primitive.ObjectID) error {
	return nil
}

func (m *MockOrderRepository) FindRestaurantAllOrders(context.Context, uint) ([]model.Order, error) {
	return []model.Order{}, nil
}

func (m *MockOrderRepository) FindRestaurantCurrentOrders(context.Context, uint) ([]model.Order, error) {
	return []model.Order{}, nil
}

func (m *MockOrderRepository) FindUserAllOrders(context.Context, uint) ([]model.Order, error) {
	return []model.Order{}, nil
}

func (m *MockOrderRepository) FindUserCurrentOrders(context.Context, uint) ([]model.Order, error) {
	return []model.Order{}, nil
}

func (m *MockOrderRepository) FindDriverAllOrders(context.Context, uint) ([]model.Order, error) {
	return []model.Order{}, nil
}

func (m *MockOrderRepository) FindDriverCurrentOrder(context.Context, uint) (model.Order, error) {
	return model.Order{}, nil
}

func (m *MockOrderRepository) UpdateRestaurantStatus(ctx context.Context, orderId primitive.ObjectID, userId uint32, status string) error {
	return nil
}

func (m *MockOrderRepository) UpdateDriverStatus(ctx context.Context, orderId primitive.ObjectID, userId uint32, status string) error {
	return nil
}

func (m *MockOrderRepository) UpdatePaymentStatus(ctx context.Context, orderId primitive.ObjectID, status, method, completedAt string) error {
	return nil
}