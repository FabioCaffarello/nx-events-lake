// go:build wireinject
// +build wireinject

package main

import (
	"apps/services-orchestration/services-file-catalog-handler/internal/entity"
	"apps/services-orchestration/services-file-catalog-handler/internal/infra/database"
	webHandler "apps/services-orchestration/services-file-catalog-handler/internal/infra/web/handlers"

	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

var setFileCatalogRepositoryDependency = wire.NewSet(
	database.NewFileCatalogRepository,
	wire.Bind(
		new(entity.FileCatalogInterface),
		new(*database.FileCatalogRepository),
	),
)

func NewWebFileCatalogHandler(client *mongo.Client, database string) *webHandler.WebFileCatalogHandler {
	wire.Build(
		setFileCatalogRepositoryDependency,
		webHandler.NewWebFileCatalogHandler,
	)
	return &webHandler.WebFileCatalogHandler{}
}

func NewHealthzHandler() *webHandler.WebHealthzHandler {
    wire.Build(
         webHandler.NewWebHealthzHandler,
    )
    return &webHandler.WebHealthzHandler{}
}
