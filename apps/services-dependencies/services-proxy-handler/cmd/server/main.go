package main

import (
	"apps/services-dependencies/services-proxy-handler/configs"
	"apps/services-dependencies/services-proxy-handler/internal/infra/web/webserver"
	"log"
	"os"
)

func main() {
	environment := os.Getenv("ENVIRONMENT")
	configs, err := configs.LoadConfig(".", environment)
	if err != nil {
		panic(err)
	}

	healthzHandler := NewHealthzHandler()
    torProxyHandler := NewTorProxyHandler()

	webserver := webserver.NewWebServer(configs.WebServerPort)
	webserver.AddHandler("/", "GET", "/tor", torProxyHandler.GetTorIPRotation)


	webserver.HandleHealthz(healthzHandler.Healthz)

	log.Printf("Server started at port %s", configs.WebServerPort)
	webserver.Start()

	select {}
}
