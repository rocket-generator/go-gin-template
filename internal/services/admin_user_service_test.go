package services

import (
	"github.com/takaaki-mizuno/go-gin-template/config"
	"github.com/takaaki-mizuno/go-gin-template/internal/repositories"
	"github.com/takaaki-mizuno/go-gin-template/pkg/database"
	"github.com/takaaki-mizuno/go-gin-template/pkg/hash"
	"github.com/takaaki-mizuno/go-gin-template/pkg/logger"
	"github.com/takaaki-mizuno/go-gin-template/pkg/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAdminUserService(t *testing.T) {
	t.Run("Create NewAdminUserService", func(t *testing.T) {
		service := createAdminUserService()
		assert.NotNil(t, service)
	})
}

func createAdminUserService() AdminUserServiceInterface {
	configInstance, _ := config.LoadConfig()
	db, _, _ := database.GetMockDB()

	service := NewAdminUserService(
		&repositories.MockAdminUserRepository{},
		configInstance,
		db,
		logger.NewTestLogger(),
		&token.ProviderMock{},
		&hash.ProviderMock{},
	)

	return service
}
