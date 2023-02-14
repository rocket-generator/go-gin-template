package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/takaaki-mizuno/go-gin-template/internal/http/admin"
	"github.com/takaaki-mizuno/go-gin-template/internal/http/app"
	"github.com/takaaki-mizuno/go-gin-template/internal/http/common"
	"go.uber.org/dig"
	"os"
)

func Bootstrap(container *dig.Container) (*gin.Engine, error) {
	var err error

	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	gin.DefaultWriter = os.Stderr

	_app := gin.New()

	err = setupMiddlewares(_app, container)
	if err != nil {
		panic(err)
	}
	err = common.SetupRoutes(_app, container)
	if err != nil {
		panic(err)
	}
	err = app.SetupRoutes(_app, container)
	if err != nil {
		panic(err)
	}
	err = admin.SetupRoutes(_app, container)

	return _app, err
}
