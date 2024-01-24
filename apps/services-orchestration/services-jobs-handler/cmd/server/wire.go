// go:build wireinject
// +build wireinject

package main

import (
	"apps/services-orchestration/services-jobs-handler/internal/entity"
	"apps/services-orchestration/services-jobs-handler/internal/infra/database"
	webHandler "apps/services-orchestration/services-jobs-handler/internal/infra/web/handlers"

	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

var setHttpGatewayParamsRepositoryDependency = wire.NewSet(
	database.NewHttpGatewayParamsRepository,
	wire.Bind(
		new(entity.HttpGatewayParamsInterface),
		new(*database.HttpGatewayParamsRepository),
	),
)

var setHttpGatewayParamsVersionRepositoryDependency = wire.NewSet(
	database.NewHttpGatewayParamsVersionRepository,
	wire.Bind(
		new(entity.HttpGatewayParamsVersionInterface),
		new(*database.HttpGatewayParamsVersionRepository),
	),
)

func NewWebJobParamsHttpGatewayHandler(client *mongo.Client, database string) *webHandler.WebJobParamsHttpGatewayHandler {
	wire.Build(
		setHttpGatewayParamsRepositoryDependency,
		setHttpGatewayParamsVersionRepositoryDependency,
		webHandler.NewWebJobParamsHttpGatewayHandler,
	)
	return &webHandler.WebJobParamsHttpGatewayHandler{}
}

func NewHealthzHandler() *webHandler.WebHealthzHandler {
	wire.Build(
		webHandler.NewWebHealthzHandler,
	)
	return &webHandler.WebHealthzHandler{}
}
