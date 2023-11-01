package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"libs/golang/shared/go-events/events"
	"libs/golang/resources/go-rabbitmq/queue"
)

type ConfigCreatedHandler struct {
	RabbitMQ *queue.RabbitMQ
}

func NewConfigCreatedHandler(rabbitMQ *queue.RabbitMQ) *ConfigCreatedHandler {
	return &ConfigCreatedHandler{
		RabbitMQ: rabbitMQ,
	}
}

func (si *ConfigCreatedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup, exchangeName string, routingKey string) {
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
