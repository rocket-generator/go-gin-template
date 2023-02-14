package cmd

import (
	"github.com/takaaki-mizuno/go-gin-template/config"
	"github.com/takaaki-mizuno/go-gin-template/pkg/database"
	"github.com/takaaki-mizuno/go-gin-template/pkg/logger"
	"go.uber.org/dig"
)

// BuildContainer ...
func BuildContainer() *dig.Container {
	container := dig.New()
	_ = container.Provide(config.LoadConfig)
	_ = container.Provide(logger.NewLogger)
	_ = container.Provide(database.InitDatabase)
	return container
}
