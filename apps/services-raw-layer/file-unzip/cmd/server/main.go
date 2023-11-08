package main

import (
	"apps/services-raw-layer/file-unzip/configs"
	"apps/services-raw-layer/file-unzip/internal/infra/consumer"
	"apps/services-raw-layer/file-unzip/internal/infra/consumer/listener"
	"fmt"
	"os"
)

const (
	inputQueueName = "inputs-unzip"
)

func main() {
	environment := os.Getenv("ENVIRONMENT")
	configs, err := configs.LoadConfig(".", environment)
	if err != nil {
		panic(err)
	}

	inputRountingKey := fmt.Sprintf("%s.%s.inputs.*", configs.ContextEnv, configs.ServiceName)

	consumers := consumer.NewConsumer(configs)
	inputListener := listener.NewServiceInputListener(configs)

	consumers.Register(inputQueueName, inputRountingKey, inputListener)
	consumers.RunConsumers()
}
