package cmd

import (
	"github.com/takaaki-mizuno/go-gin-template/config"
	adminMiddlewares "github.com/takaaki-mizuno/go-gin-template/internal/http/admin/middlewares"
	commonHandlers "github.com/takaaki-mizuno/go-gin-template/internal/http/common/handlers"
	"github.com/takaaki-mizuno/go-gin-template/internal/repositories"
	"github.com/takaaki-mizuno/go-gin-template/internal/services"
	"github.com/takaaki-mizuno/go-gin-template/pkg/database"
	"github.com/takaaki-mizuno/go-gin-template/pkg/hash"
	"github.com/takaaki-mizuno/go-gin-template/pkg/logger"
	commonMiddlewares "github.com/takaaki-mizuno/go-gin-template/pkg/middlewares"
	"github.com/takaaki-mizuno/go-gin-template/pkg/token"
	"go.uber.org/dig"
)

// BuildContainer ...
func BuildContainer() *dig.Container {
	container := dig.New()
	_ = container.Provide(config.LoadConfig)
	_ = container.Provide(logger.NewLogger)
	_ = container.Provide(database.InitDatabase)

	// Packages
	_ = container.Provide(
		func(config *config.Config) (token.ProviderInterface, error) {
			return token.NewTokenProvider([]byte(config.Authentication.Secret))
		},
	)
	_ = container.Provide(hash.NewHashProvider)

	// Middlewares
	_ = container.Provide(commonMiddlewares.SecurityHeaders, dig.Name("securityHeaders"))
	_ = container.Provide(commonMiddlewares.RequestID, dig.Name("requestID"))
	_ = container.Provide(commonMiddlewares.Logger, dig.Name("logger"))
	_ = container.Provide(adminMiddlewares.Auth, dig.Name("adminAuth"))

	// Repositories
	_ = container.Provide(repositories.NewAdminUserRepository)

	// Handlers
	_ = container.Provide(commonHandlers.NewHandler)

	// Services
	_ = container.Provide(services.NewAdminUserService)

	return container
}
