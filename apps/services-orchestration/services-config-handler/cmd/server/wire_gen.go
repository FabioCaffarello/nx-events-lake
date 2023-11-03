// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"apps/services-orchestration/services-config-handler/internal/entity"
	"apps/services-orchestration/services-config-handler/internal/event"
	"apps/services-orchestration/services-config-handler/internal/infra/database"
	"apps/services-orchestration/services-config-handler/internal/infra/web/handlers"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"libs/golang/shared/go-events/events"
)

// Injectors from wire.go:

func NewWebConfigHandler(client *mongo.Client, eventDispatcher events.EventDispatcherInterface, database2 string) *handlers.WebConfigHandler {
	configRepository := database.NewConfigRepository(client, database2)
	configCreated := event.NewConfigCreated()
	configVersionRepository := database.NewConfigVersionRepository(client, database2)
	webConfigHandler := handlers.NewWebConfigHandler(eventDispatcher, configRepository, configCreated, configVersionRepository)
	return webConfigHandler
}

func NewHealthzHandler() *handlers.WebHealthzHandler {
	webHealthzHandler := handlers.NewWebHealthzHandler()
	return webHealthzHandler
}

// wire.go:

var setConfigRepositoryDependency = wire.NewSet(database.NewConfigRepository, wire.Bind(
	new(entity.ConfigInterface),
	new(*database.ConfigRepository),
),
)

var setConfigVersionRepositoryDependency = wire.NewSet(database.NewConfigVersionRepository, wire.Bind(
	new(entity.ConfigVersionInterface),
	new(*database.ConfigVersionRepository),
),
)

var setConfigCreatedEvent = wire.NewSet(event.NewConfigCreated, wire.Bind(new(events.EventInterface), new(*event.ConfigCreated)))