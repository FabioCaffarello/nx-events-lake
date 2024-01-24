package main

import (
	"apps/services-orchestration/services-staging-handler/configs"
	"apps/services-orchestration/services-staging-handler/internal/event/handler"
	"apps/services-orchestration/services-staging-handler/internal/infra/web/webserver"
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
	eventDispatcher.Register("ProcessingJobDependenciesCreated", &handler.ProcessingJobDependenciesCreatedHandler{
		RabbitMQ: rabbitMQ,
	})

	healthzUseCase := NewHealthzHandler()

	// Web
	webserver := webserver.NewWebServer(configs.WebServerPort)

	webProcessingJobDependenciesHandler := NewWebProcessingJobDependenciesHandler(client, eventDispatcher, configs.DBName)
    webProcessingGraphHandler := NewWebProcessingGraphsHandler(client, configs.DBName)

    webserver.AddHandler("/", "POST", "/jobs-dependencies", webProcessingJobDependenciesHandler.CreateProcessingJobDependenciesHandler)
    webserver.AddHandler("/", "GET", "/jobs-dependencies/{id}", webProcessingJobDependenciesHandler.ListOneProcessingJobDependenciesByIdHandler)
    webserver.AddHandler("/", "DELETE", "/jobs-dependencies/{id}", webProcessingJobDependenciesHandler.RemoveProcessingJobDependenciesHandler)
    webserver.AddHandler("/", "POST", "/jobs-dependencies/{id}", webProcessingJobDependenciesHandler.UpdateProcessingJobDependenciesHandler)

    webserver.AddHandler("/", "POST", "/processing-graph", webProcessingGraphHandler.CreateProcessingGraphHandler)
    webserver.AddHandler("/", "GET", "/processing-graph/source/{source}/start-processing-id/{start_processing_id}", webProcessingGraphHandler.ListOneProcessingGraphBySourceAndStartProcessingIdUseCase)
    webserver.AddHandler("/", "POST", "/processing-graph/source/{source}/start-processing-id/{start_processing_id}", webProcessingGraphHandler.CreateTaskToProcessingGraphHandler)
    webserver.AddHandler("/", "GET", "/processing-graph/source/{source}/parent-processing-id/{parent_processing_id}", webProcessingGraphHandler.ListOneProcessingGraphByTaskSourceAndParentProcessingIdUseCase)
    webserver.AddHandler("/", "POST", "/processing-graph/source/{source}/processing-id/{processing_id}/status/{status}/processing-date/{processing_timestamp}", webProcessingGraphHandler.UpdateTaskStatusProcessingGraphHandler)
    webserver.AddHandler("/", "POST", "/processing-graph/source/{source}/processing-id/{processing_id}/output-id/{output_id}", webProcessingGraphHandler.UpdateTaskOutputProcessingGraphHandler)

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
