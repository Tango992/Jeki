package service

import (
	"email-service/helper"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type MailService struct {
	channel *amqp.Channel
}

func NewMailService(ch *amqp.Channel) Mail {
	return MailService{
		channel: ch,
	}
}

func (m MailService) SendVerificationEmail(q amqp.Queue) {
	msgs, err := m.channel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatal(err)
	}

	for d := range msgs {
		log.Printf("\033[36mNEW MESSAGE:\033[0m %s", d.Body)

		userData := helper.AssertJsonToStruct(d.Body)
		if err := helper.SendVerificationEmail(userData); err != nil {
			log.Fatal(err)
		}
	}
}