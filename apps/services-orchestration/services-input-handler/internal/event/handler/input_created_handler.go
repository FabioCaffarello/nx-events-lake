package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"libs/golang/resources/go-rabbitmq/queue"
	"libs/golang/shared/go-events/events"
)

type InputCreatedHandler struct {
	RabbitMQ *queue.RabbitMQ
}

func NewInputCreatedHandler(rabbitMQ *queue.RabbitMQ) *InputCreatedHandler {
	return &InputCreatedHandler{
		RabbitMQ: rabbitMQ,
	}
}

func (si *InputCreatedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup, exchangeName string, routingKey string) {
	defer wg.Done()
	jsonOutput, _ := json.Marshal(event.GetPayload())
	err := si.RabbitMQ.Notify(
		jsonOutput,
		"application/json",
		exchangeName,
		routingKey,
	)
	if err != nil {
		fmt.Println(err)
	}
}
