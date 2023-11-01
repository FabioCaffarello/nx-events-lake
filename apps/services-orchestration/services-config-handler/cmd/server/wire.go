// go:build wireinject
// +build wireinject

package main

import (
	"apps/services-orchestration/services-config-handler/internal/entity"
	"apps/services-orchestration/services-config-handler/internal/event"
	"apps/services-orchestration/services-config-handler/internal/infra/database"
	webHandler "apps/services-orchestration/services-config-handler/internal/infra/web/handlers"
	"libs/golang/shared/go-events/events"

	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

var setConfigRepositoryDependency = wire.NewSet(
	database.NewConfigRepository,
	wire.Bind(
		new(entity.ConfigInterface),
		new(*database.ConfigRepository),
	),
)

var setConfigVersionRepositoryDependency = wire.NewSet(
	database.NewConfigVersionRepository,
	wire.Bind(
		new(entity.ConfigVersionInterface),
		new(*database.ConfigVersionRepository),
	),
)

var setConfigCreatedEvent = wire.NewSet(
	event.NewConfigCreated,
	wire.Bind(new(events.EventInterface), new(*event.ConfigCreated)),
)

func NewWebConfigHandler(client *mongo.Client, eventDispatcher events.EventDispatcherInterface, database string) *webHandler.WebConfigHandler {
	wire.Build(
		setConfigRepositoryDependency,
		setConfigVersionRepositoryDependency,
		setConfigCreatedEvent,
		webHandler.NewWebConfigHandler,
	)
	return &webHandler.WebConfigHandler{}
}


func NewHealthzHandler() *webHandler.WebHealthzHandler {
    wire.Build(
         webHandler.NewWebHealthzHandler,
    )
    return &webHandler.WebHealthzHandler{}
}
