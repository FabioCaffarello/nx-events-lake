package main

import (
	"apps/services-orchestration/services-file-catalog-handler/configs"
	"apps/services-orchestration/services-file-catalog-handler/internal/infra/web/webserver"
	"context"
	"fmt"
	mongoClient "libs/golang/resources/go-mongo/client"
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

	// Web
	webserver := webserver.NewWebServer(configs.WebServerPort)

	webFileCatalogHandler := NewWebFileCatalogHandler(client, configs.DBName)

    webserver.AddHandler("/", "POST", "/file-catalog", webFileCatalogHandler.CreateFileCatalog)
    webserver.AddHandler("/", "GET", "/file-catalog", webFileCatalogHandler.ListAllFileCatalogs)
    webserver.AddHandler("/", "GET", "/file-catalog/{id}", webFileCatalogHandler.ListOneFileCatalogById)
    webserver.AddHandler("/", "GET", "/file-catalog/service/{service}", webFileCatalogHandler.ListAllFileCatalogsByService)
    webserver.AddHandler("/", "GET", "/file-catalog/service/{service}/source/{source}", webFileCatalogHandler.ListOneFileCatalogByServiceAndSource)

	webserver.HandleHealthz(healthzUseCase.Healthz)

	fmt.Println("Server is running on port", configs.WebServerPort)
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
