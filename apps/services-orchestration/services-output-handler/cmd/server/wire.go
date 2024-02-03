// go:build wireinject
//go:build wireinject
// +build wireinject

package main

import (
	"apps/services-orchestration/services-output-handler/internal/entity"
	"apps/services-orchestration/services-output-handler/internal/infra/database"
	webHandler "apps/services-orchestration/services-output-handler/internal/infra/web/handlers"

	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

var setServiceOutputRepositoryDependency = wire.NewSet(
	database.NewServiceOutputRepository,
	wire.Bind(
		new(entity.ServiceOutputInterface),
		new(*database.ServiceOutputRepository),
	),
)

func NewWebServiceOutputHandler(client *mongo.Client, database string) *webHandler.WebServiceOutputHandler {
	wire.Build(
		setServiceOutputRepositoryDependency,
		webHandler.NewWebServiceOutputHandler,
	)
	return &webHandler.WebServiceOutputHandler{}
}

func NewHealthzHandler() *webHandler.WebHealthzHandler {
	wire.Build(
		webHandler.NewWebHealthzHandler,
	)
	return &webHandler.WebHealthzHandler{}
}
