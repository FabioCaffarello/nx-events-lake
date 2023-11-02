//go:build wireinject
// +build wireinject

package main

import (
	"apps/services-orchestration/services-input-handler/internal/entity"
	"apps/services-orchestration/services-input-handler/internal/event"
	"apps/services-orchestration/services-input-handler/internal/infra/database"
	webHandler "apps/services-orchestration/services-input-handler/internal/infra/web/handlers"
	"libs/golang/shared/go-events/events"

	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

var setInputRepositoryDependency = wire.NewSet(
	database.NewInputRepository,
	wire.Bind(
		new(entity.InputInterface),
		new(*database.InputRepository),
	),
)

var setInputCreatedEvent = wire.NewSet(
	event.NewInputCreated,
	wire.Bind(new(events.EventInterface), new(*event.InputCreated)),
)

var setInputStatusUpdatedEvent = wire.NewSet(
	event.NewInputStatusUpdated,
	wire.Bind(new(events.EventInterface), new(*event.InputStatusUpdated)),
)

func NewWebInputHandler(client *mongo.Client, eventDispatcher events.EventDispatcherInterface, database string) *webHandler.WebInputHandler {
	wire.Build(
		setInputRepositoryDependency,
		setInputCreatedEvent,
		webHandler.NewWebInputHandler,
	)
	return &webHandler.WebInputHandler{}
}

func NewWebInputStatusHandler(client *mongo.Client, eventDispatcher events.EventDispatcherInterface, database string) *webHandler.WebInputStatusHandler {
	wire.Build(
		setInputRepositoryDependency,
		setInputStatusUpdatedEvent,
		webHandler.NewWebInputStatusHandler,
	)
	return &webHandler.WebInputStatusHandler{}
}

func NewHealthzHandler() *webHandler.WebHealthzHandler {
    wire.Build(
         webHandler.NewWebHealthzHandler,
    )
    return &webHandler.WebHealthzHandler{}
}
