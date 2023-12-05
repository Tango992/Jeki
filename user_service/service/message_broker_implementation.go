package service

import (
	"context"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RabbitMQ struct {
	Ch *amqp.Channel
}

func NewMessageBroker(ch *amqp.Channel) MessageBroker {
	return RabbitMQ{
		Ch: ch,
	}
}

func (r RabbitMQ) PublishMessage(message []byte) error {
	q, err := r.Ch.QueueDeclare(
		"email_verification", 	// name
		false,		// durable
		false,		// delete when unused
		false,		// exclusive
		false,		// no-wait
		nil,		// arguments
	)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	ctx := context.TODO()

	err = r.Ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		})
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	log.Printf(" [x] Sent %s\n", message)
	return nil
}
