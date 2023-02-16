package cmd

import (
	"github.com/takaaki-mizuno/go-gin-template/config"
	adminMiddlewares "github.com/takaaki-mizuno/go-gin-template/internal/http/admin/middlewares"
	"github.com/takaaki-mizuno/go-gin-template/internal/repositories"
	"github.com/takaaki-mizuno/go-gin-template/internal/services"
	"github.com/takaaki-mizuno/go-gin-template/pkg/database"
	"github.com/takaaki-mizuno/go-gin-template/pkg/hash"
	"github.com/takaaki-mizuno/go-gin-template/pkg/logger"
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
	_ = container.Provide(adminMiddlewares.Auth, dig.Name("adminAuth"))

	// Repositories
	_ = container.Provide(repositories.NewAdminUserRepository)

	// Services
	_ = container.Provide(services.NewAdminUserService)

	return container
}
