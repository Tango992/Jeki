package service

type MessageBroker interface {
	PublishMessage(message []byte) error
}