// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"fast-wire/app/config"
)

// Injectors from wire.go:

func InitializeApplication(config2 config.Config) (application, error) {
	handler := provideRouter()
	server := provideServer(handler, config2)
	mainApplication := newApplication(server)
	return mainApplication, nil
}