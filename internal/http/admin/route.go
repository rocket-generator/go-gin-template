package admin

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func SetupRoutes(app *gin.Engine, container *dig.Container) error {
	/*
		var _handlers handlers.HandlerInterface
		if err := container.Invoke(func(
			h handlers.HandlerInterface,
		) {
			_handlers = h
		}); err != nil {
			return err
		}
	*/
	return nil
}
