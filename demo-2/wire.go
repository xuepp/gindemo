//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitApp() (*gin.Engine, *Config, error) {
	wire.Build(
		NewConfig,
		NewDB,
		NewUserRepository,
		NewUserService,
		NewUserHandler,
		NewGinEngine,
	)
	return nil, nil, nil
}

// NewGinEngine 用于组装 Gin
func NewGinEngine(handler *UserHandler, cfg *Config) *gin.Engine {
	e := gin.Default()
	handler.RegisterRoutes(e)
	return e
}
