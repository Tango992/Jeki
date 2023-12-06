package service

import (
	"order-service/model"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockPaymentService struct {
	Mock mock.Mock
}

func NewMockPaymentService() MockPaymentService {
	return MockPaymentService{}
}

func (m *MockPaymentService) MakeInvoice(externalID primitive.ObjectID, subtotal float32) (model.Payment, error) {
	args := m.Mock.Called(externalID, subtotal)
	return args.Get(0).(model.Payment), args.Error(1)
}