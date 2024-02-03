// go:build wireinject
// +build wireinject

package main

import (
	webHandler "apps/services-dependencies/services-proxy-handler/internal/infra/web/handlers"

	"github.com/google/wire"
)

func NewHealthzHandler() *webHandler.WebHealthzHandler {
	wire.Build(
		webHandler.NewWebHealthzHandler,
	)
	return &webHandler.WebHealthzHandler{}
}


func NewTorProxyHandler() *webHandler.TorProxyHandler {
	wire.Build(
		webHandler.NewTorProxyHandler,
	)
	return &webHandler.TorProxyHandler{}
}
