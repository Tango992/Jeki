package config

import (
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func InitMbChannel() (*amqp.Connection, *amqp.Channel){
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

func InitMbQueue(ch *amqp.Channel) amqp.Queue {
	q, err := ch.QueueDeclare(
		"email_verification", // name
		false,  // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	if err != nil {
		log.Fatal(err)
	}
	return q
}