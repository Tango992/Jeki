package config

import (
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func InitMessageBroker() (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(os.Getenv("RABBIT_MQ_URL"))
	if err != nil {
		log.Fatal(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	return conn, ch
}
