package main

import (
	"apps/services-orchestration/services-output-handler/configs"
	"apps/services-orchestration/services-output-handler/internal/infra/web/webserver"
	"context"
	mongoClient "libs/golang/resources/go-mongo/client"
	"log"
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

    healthzUseCase := NewHealthzHandler()

	webserver := webserver.NewWebServer(configs.WebServerPort)
	webServiceOutputHandler := NewWebServiceOutputHandler(client, configs.DBName)

	webserver.AddHandler("/outputs", "POST", "/outputs/service/{service}", webServiceOutputHandler.CreateServiceOutput)
	webserver.AddHandler("/outputs", "GET", "/outputs/service/{service}/{id}", webServiceOutputHandler.ListOneServiceOutputsByServiceAndId)
	webserver.AddHandler("/outputs", "GET", "/outputs/service/{service}", webServiceOutputHandler.ListAllServiceOutputsByService)
	webserver.AddHandler("/outputs", "GET", "/outputs/service/{service}/source/{source}", webServiceOutputHandler.ListAllServiceOutputsByServiceAndSource)

    webserver.HandleHealthz(healthzUseCase.Healthz)

	log.Printf("Server started at port %s", configs.WebServerPort)
	webserver.Start()

	select {}
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
