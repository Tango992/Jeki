package config

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func InitMessageBroker() (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	return conn, ch
}
