package serve

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/takaaki-mizuno/go-gin-template/cmd"
	"github.com/takaaki-mizuno/go-gin-template/config"
	"github.com/takaaki-mizuno/go-gin-template/internal/http/common"
	"go.uber.org/dig"
	"go.uber.org/zap"
	"os"
	"strconv"
)

// NewAppServer ... run the app server
func NewAppServer(command *cobra.Command, args []string) error {
	var configInstance *config.Config
	var logger *zap.Logger

	container := cmd.BuildContainer()

	app, err := bootstrap(container)
	if err != nil {
		return err
	}

	err = container.Invoke(func(_config *config.Config, _logger *zap.Logger) {
		configInstance = _config
		logger = _logger
	})
	if err != nil {
		return err
	}

	logger.Info("starting APP server ( APIs for Users )",
		zap.String("url", configInstance.App.APP.URL),
		zap.Uint("port", configInstance.App.APP.Port),
	)

	err = app.Run(":" + strconv.FormatUint(uint64(configInstance.App.APP.Port), 10))

	return err
}

func bootstrap(container *dig.Container) (*gin.Engine, error) {
	var err error

	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	gin.DefaultWriter = os.Stderr

	app := gin.New()

	err = setupMiddlewares(app, container)
	if err == nil {
		err = common.SetupRoutes(app, container)

	}

	return app, err
}
