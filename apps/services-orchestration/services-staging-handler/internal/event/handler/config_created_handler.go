package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"libs/golang/shared/go-events/events"
	"libs/golang/resources/go-rabbitmq/queue"
)

type ProcessingJobDependenciesCreatedHandler struct {
	RabbitMQ *queue.RabbitMQ
}

func NewProcessingJobDependenciesCreatedHandler(rabbitMQ *queue.RabbitMQ) *ProcessingJobDependenciesCreatedHandler {
	return &ProcessingJobDependenciesCreatedHandler{
		RabbitMQ: rabbitMQ,
	}
}

func (si *ProcessingJobDependenciesCreatedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup, exchangeName string, routingKey string) {
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
