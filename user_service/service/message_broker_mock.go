package service

import (
	"github.com/stretchr/testify/mock"
)

type MockMessageBroker struct {
	Mock mock.Mock
}

func NewMockMessageBroker() MockMessageBroker {
	return MockMessageBroker{}
}

func (m *MockMessageBroker) PublishMessage(message []byte) error {
	args := m.Mock.Called(message)
	return args.Error(0)
}