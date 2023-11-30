package service

import amqp "github.com/rabbitmq/amqp091-go"

type Mail interface {
	SendVerificationEmail(amqp.Queue)
}