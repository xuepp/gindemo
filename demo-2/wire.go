//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitApp() (*gin.Engine, error) {
	wire.Build(
		NewConfig,
		NewDB,
		NewUserRepository,
		NewUserService,
		NewUserHandler,
		NewGinEngine,
	)
	return nil, nil
}

func NewGinEngine(handler *UserHandler, cfg *Config) *gin.Engine {
	r := gin.Default()
	handler.RegisterRoutes(r)
	return r
}
