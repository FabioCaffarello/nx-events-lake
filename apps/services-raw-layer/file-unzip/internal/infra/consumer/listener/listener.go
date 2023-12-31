package listener

import (
	"libs/golang/resources/go-rabbitmq/queue"

	amqp "github.com/rabbitmq/amqp091-go"
)

type MessageHandlerInterface interface {
	Handle(rabbitMQ *queue.RabbitMQ, exchange string, msg amqp.Delivery) error
}
