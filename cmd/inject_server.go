package main

import (
	"fast-wire/app/config"
	"fast-wire/app/handler/web"
	"fast-wire/app/server"

	"github.com/fate-lovely/phi"
	"github.com/google/wire"
)

var serverSet = wire.NewSet(
	provideRouter,
	provideServer,
)

func provideServer(handler phi.Handler, config config.Config) *server.Server {
	return &server.Server{
		Host:    config.Server.Host,
		Handler: handler,
	}
}

func provideRouter() phi.Handler {
	r := phi.NewRouter()

	r.Mount("/", web.Handler())
	return r
}
