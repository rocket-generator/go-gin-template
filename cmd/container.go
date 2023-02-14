package cmd

import (
	"github.com/takaaki-mizuno/go-gin-template/config"
	"go.uber.org/dig"
)

// BuildContainer ...
func BuildContainer() *dig.Container {
	container := dig.New()
	_ = container.Provide(config.LoadConfig)
	return container
}
