package common

import (
	"github.com/gin-gonic/gin"
	"github.com/takaaki-mizuno/go-gin-template/internal/http/common/handlers"
	"go.uber.org/dig"
)

func SetupRoutes(app *gin.Engine, container *dig.Container) error {
	var _handlers handlers.HandlerInterface
	if err := container.Invoke(func(
		h handlers.HandlerInterface,
	) {
		_handlers = h
	}); err != nil {
		return err
	}

	app.NoRoute(_handlers.NoRoute)
	app.NoMethod(_handlers.NoMethod)

	app.GET("/", _handlers.IndexGet)
	app.GET("/healthz", _handlers.Healthz)

	return nil
}
