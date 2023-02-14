package config

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/takaaki-mizuno/go-gin-template/cmd"
	"github.com/takaaki-mizuno/go-gin-template/config"
	"go.uber.org/zap"
)

func Print(command *cobra.Command, args []string) error {
	var logger *zap.Logger
	var configInstance *config.Config

	container := cmd.BuildContainer()
	err := container.Invoke(func(_logger *zap.Logger, _config *config.Config) {
		logger = _logger
		configInstance = _config
	})

	b, err := json.MarshalIndent(configInstance, "", "    ")
	if err != nil {
		logger.Fatal("failed to marshal config", zap.Error(err))
		return err
	}
	fmt.Printf("config: %s", string(b))
	return nil
}
