package main

import (
	"apps/services-orchestration/services-config-handler/configs"
	"apps/services-orchestration/services-config-handler/internal/event/handler"
	"apps/services-orchestration/services-config-handler/internal/infra/web/webserver"
	"context"
	"fmt"
	mongoClient "libs/golang/resources/go-mongo/client"
	"libs/golang/resources/go-rabbitmq/queue"
	"libs/golang/shared/go-events/events"
	"os"
	"time"
)

func main() {
	environment := os.Getenv("ENVIRONMENT")
	configs, err := configs.LoadConfig(".", environment)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	mongoDB := getMongoDBClient(configs, ctx)
	client := mongoDB.Client
	defer client.Disconnect(ctx)

	rabbitMQ := getRabbitMQChannel(configs)
	defer rabbitMQ.Close()

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("ConfigCreated", &handler.ConfigCreatedHandler{
		RabbitMQ: rabbitMQ,
	})

	healthzUseCase := NewHealthzHandler()

	// Web
	webserver := webserver.NewWebServer(configs.WebServerPort)

	webConfigHandler := NewWebConfigHandler(client, eventDispatcher, configs.DBName)

	webserver.AddHandler("/", "POST", "/configs", webConfigHandler.CreateConfig)
	webserver.AddHandler("/", "POST", "/configs/update", webConfigHandler.UpdateConfig)
	webserver.AddHandler("/", "GET", "/configs", webConfigHandler.ListAllConfigs)
	webserver.AddHandler("/", "GET", "/configs/versions", webConfigHandler.ListAllConfigVersion)
	webserver.AddHandler("/", "GET", "/configs/{id}", webConfigHandler.ListOneConfigById)
	webserver.AddHandler("/", "GET", "/configs/versions/{id}", webConfigHandler.ListOneConfigVersionById)
	webserver.AddHandler("/", "GET", "/configs/versions/{id}/version-id/{versionId}", webConfigHandler.ListOneConfigByIdAndVersionId)
	webserver.AddHandler("/", "GET", "/configs/service/{service}", webConfigHandler.ListAllConfigsByService)
	webserver.AddHandler("/", "GET", "/configs/service/{service}/context/{context}", webConfigHandler.ListAllConfigsByServiceAndContext)
	webserver.AddHandler("/", "GET", "/configs/service/{service}/source/{source}", webConfigHandler.ListAllConfigsByDependentJob)
	webserver.AddHandler("/", "GET", "/configs/service/{service}/source/{source}/context/{context}", webConfigHandler.ListOneConfigByServiceAndSourceAndContext)

	webserver.HandleHealthz(healthzUseCase.Healthz)

	fmt.Println("Server is running on port", configs.WebServerPort)
	webserver.Start()

	select {}
}

func getRabbitMQChannel(config configs.Config) *queue.RabbitMQ {
	rabbitMQ := queue.NewRabbitMQ(
		config.RabbitMQUser,
		config.RabbitMQPassword,
		config.RabbitMQHost,
		config.RabbitMQPort,
		config.RabbitMQVhost,
		config.RabbitMQConsumerQueueName,
		config.RabbitMQConsumerName,
		config.RabbitMQDlxName,
		config.RabbitMQProtocol,
	)
	_, err := rabbitMQ.Connect()
	if err != nil {
		panic(err)
	}
	rabbitMQ.DeclareExchange(config.RabbitMQExchange, config.RabbitMQExchangeType)
	return rabbitMQ
}

func getMongoDBClient(configs configs.Config, ctx context.Context) *mongoClient.MongoDB {
	mongoDB := mongoClient.NewMongoDB(
		configs.DBDriver,
		configs.DBUser,
		configs.DBPassword,
		configs.DBHost,
		configs.DBPort,
		configs.DBName,
		ctx,
	)

	_, err := mongoDB.Connect()
	if err != nil {
		panic(err)
	}

	return mongoDB
}
