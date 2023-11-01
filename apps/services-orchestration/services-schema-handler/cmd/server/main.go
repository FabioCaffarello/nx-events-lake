package main

import (
	"apps/services-orchestration/services-schema-handler/configs"
	"apps/services-orchestration/services-schema-handler/internal/event/handler"
	"apps/services-orchestration/services-schema-handler/internal/infra/web/webserver"
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
	eventDispatcher.Register("SchemaCreated", &handler.SchemaCreatedHandler{
		RabbitMQ: rabbitMQ,
	})

	healthzUseCase := NewHealthzHandler()

	// Web
	webserver := webserver.NewWebServer(configs.WebServerPort)

	webSchemaHandler := NewWebSchemaHandler(client, eventDispatcher, configs.DBName)

	webserver.AddHandler("/", "POST", "/schemas", webSchemaHandler.CreateSchema)
	webserver.AddHandler("/", "POST", "/schemas/update", webSchemaHandler.UpdateSchema)
	webserver.AddHandler("/", "GET", "/schemas", webSchemaHandler.ListAllSchemas)
	webserver.AddHandler("/", "GET", "/schemas/versions", webSchemaHandler.ListAllSchemaVersion)
	webserver.AddHandler("/", "GET", "/schemas/{id}", webSchemaHandler.ListOneSchemaById)
	webserver.AddHandler("/", "GET", "/schemas/versions/{id}", webSchemaHandler.ListOneSchemaVersionById)
	webserver.AddHandler("/", "GET", "/schemas/versions/{id}/version-id/{versionId}", webSchemaHandler.ListOneSchemaByIdAndVersionId)
	webserver.AddHandler("/", "GET", "/schemas/service/{service}", webSchemaHandler.ListAllSchemasByService)
	webserver.AddHandler("/", "GET", "/schemas/service/{service}/context/{context}", webSchemaHandler.ListAllSchemasByServiceAndContext)
	webserver.AddHandler("/", "GET", "/schemas/service/{service}/source/{source}/schema-type/{schemaType}", webSchemaHandler.ListOneSchemaByServiceSourceAndSchemaType)
	webserver.AddHandler("/", "GET", "/schemas/service/{service}/source/{source}/context/{context}/schema-type/{schemaType}", webSchemaHandler.ListOneSchemaByServiceAndSourceAndContextAndSchemaType)

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
