package service_test

import (
	"testing"
	"user-service/service"

	"github.com/stretchr/testify/assert"
)

var (
	messageBrokerService = service.NewMockMessageBroker()
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestPublishMessage(t *testing.T) {
	request := `{"message": "Hello world"}`

	messageBrokerService.Mock.On("PublishMessage", []byte(request)).Return(nil)
	
	err := messageBrokerService.PublishMessage([]byte(request))
	assert.Nil(t, err)
}