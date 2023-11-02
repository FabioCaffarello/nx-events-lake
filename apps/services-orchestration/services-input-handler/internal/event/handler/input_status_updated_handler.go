package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"libs/golang/resources/go-rabbitmq/queue"
	"libs/golang/shared/go-events/events"
)

type InputStatusUpdatedHandler struct {
	RabbitMQ *queue.RabbitMQ
}

func NewInputStatusUpdatedHandler(rabbitMQ *queue.RabbitMQ) *InputStatusUpdatedHandler {
	return &InputStatusUpdatedHandler{
		RabbitMQ: rabbitMQ,
	}
}

func (si *InputStatusUpdatedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup, exchangeName string, routingKey string) {
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
