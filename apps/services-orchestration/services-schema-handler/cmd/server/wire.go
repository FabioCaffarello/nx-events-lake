// go:build wireinject
// +build wireinject

package main

import (
	"apps/services-orchestration/services-schema-handler/internal/entity"
	"apps/services-orchestration/services-schema-handler/internal/event"
	"apps/services-orchestration/services-schema-handler/internal/infra/database"
	webHandler "apps/services-orchestration/services-schema-handler/internal/infra/web/handlers"
	"libs/golang/shared/go-events/events"

	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

var setSchemaRepositoryDependency = wire.NewSet(
	database.NewSchemaRepository,
	wire.Bind(
		new(entity.SchemaInterface),
		new(*database.SchemaRepository),
	),
)

var setSchemaVersionRepositoryDependency = wire.NewSet(
	database.NewSchemaVersionRepository,
	wire.Bind(
		new(entity.SchemaVersionInterface),
		new(*database.SchemaVersionRepository),
	),
)

var setSchemaCreatedEvent = wire.NewSet(
	event.NewSchemaCreated,
	wire.Bind(new(events.EventInterface), new(*event.SchemaCreated)),
)

func NewWebSchemaHandler(client *mongo.Client, eventDispatcher events.EventDispatcherInterface, database string) *webHandler.WebSchemaHandler {
	wire.Build(
		setSchemaRepositoryDependency,
		setSchemaVersionRepositoryDependency,
		setSchemaCreatedEvent,
		webHandler.NewWebSchemaHandler,
	)
	return &webHandler.WebSchemaHandler{}
}


func NewHealthzHandler() *webHandler.WebHealthzHandler {
    wire.Build(
         webHandler.NewWebHealthzHandler,
    )
    return &webHandler.WebHealthzHandler{}
}
