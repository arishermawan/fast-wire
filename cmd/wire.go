//+build wireinject

package main

import (
	"fast-wire/app/config"

	"github.com/google/wire"
)

func InitializeApplication(config config.Config) (application, error) {
	wire.Build(
		serverSet,
		newApplication,
	)
	return application{}, nil
}
