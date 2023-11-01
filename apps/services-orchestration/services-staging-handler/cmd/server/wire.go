// go:build wireinject
// +build wireinject

package main

import (
	"apps/services-orchestration/services-staging-handler/internal/entity"
	"apps/services-orchestration/services-staging-handler/internal/event"
	"apps/services-orchestration/services-staging-handler/internal/infra/database"
	webHandler "apps/services-orchestration/services-staging-handler/internal/infra/web/handlers"
	"libs/golang/shared/go-events/events"

	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

var setProcessingJobDependenciesRepositoryDependency = wire.NewSet(
	database.NewProcessingJobDependenciesRepository,
	wire.Bind(
		new(entity.ProcessingJobDependenciesInterface),
		new(*database.ProcessingJobDependenciesRepository),
	),
)

var setProcessingJobDependenciesCreatedEvent = wire.NewSet(
	event.NewProcessingJobDependenciesCreated,
	wire.Bind(new(events.EventInterface), new(*event.ProcessingJobDependenciesCreated)),
)

func NewWebProcessingJobDependenciesHandler(client *mongo.Client, eventDispatcher events.EventDispatcherInterface, database string) *webHandler.WebProcessingJobDependenciesHandler {
	wire.Build(
		setProcessingJobDependenciesRepositoryDependency,
		setProcessingJobDependenciesCreatedEvent,
		webHandler.NewWebProcessingJobDependenciesHandler,
	)
	return &webHandler.WebProcessingJobDependenciesHandler{}
}

func NewHealthzHandler() *webHandler.WebHealthzHandler {
	wire.Build(
		webHandler.NewWebHealthzHandler,
	)
	return &webHandler.WebHealthzHandler{}
}
