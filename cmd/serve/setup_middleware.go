package serve

import (
	"github.com/gin-gonic/gin"
	commonMiddlewares "github.com/takaaki-mizuno/go-gin-template/pkg/middlewares"
	"go.uber.org/dig"
)

func setupMiddlewares(app *gin.Engine, container *dig.Container) error {
	var securityHeadersMiddleware gin.HandlerFunc
	var requestIDMiddleware gin.HandlerFunc
	var loggerMiddleware gin.HandlerFunc

	if err := container.Invoke(func(
		_commonMiddlewares commonMiddlewares.Middlewares,
	) {
		securityHeadersMiddleware = _commonMiddlewares.SecurityHeaders
		requestIDMiddleware = _commonMiddlewares.RequestID
		loggerMiddleware = _commonMiddlewares.Logger
	}); err != nil {
		return err
	}

	app.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, ResponseType, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	app.Use(gin.Recovery())
	app.Use(securityHeadersMiddleware)
	app.Use(requestIDMiddleware)
	app.Use(loggerMiddleware)

	return nil
}
