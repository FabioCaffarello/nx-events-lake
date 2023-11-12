package main

import (
	"apps/services-orchestration/services-events-handler/configs"
	"apps/services-orchestration/services-events-handler/internal/infra/listener"
	"apps/services-orchestration/services-events-handler/internal/infra/listener/channels"
	"os"
)

const (
	feedbackQueueName = "service-feedback"
	inputQueueName    = "input-process"

	inputRountingKey   = "input-processing"
	feedbackRoutingKey = "feedback"
)

func main() {
	environment := os.Getenv("ENVIRONMENT")
	configs, err := configs.LoadConfig(".", environment)
	if err != nil {
		panic(err)
	}

	consumers := listener.NewConsumer(configs)
	serviceFeedbackListener := channels.NewServiceFeedbackListener()
	inputListener := channels.NewServiceInputListener()

	consumers.Register(inputQueueName, inputRountingKey, inputListener)
	consumers.Register(feedbackQueueName, feedbackRoutingKey, serviceFeedbackListener)

	consumers.RunConsumers()

}
