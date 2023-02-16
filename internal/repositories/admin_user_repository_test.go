package repositories

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/takaaki-mizuno/go-gin-template/pkg/database"
	"github.com/takaaki-mizuno/go-gin-template/pkg/logger"
	"testing"
)

func createAdminUserRepositoryWithMockDB() (AdminUserRepositoryInterface, sqlmock.Sqlmock) {
	db, mock, _ := database.GetMockDB()
	repository := NewAdminUserRepository(logger.NewTestLogger(), db)
	return repository, mock
}

func TestNewAdminUserRepository(t *testing.T) {
	t.Run("Create AdminUserRepository", func(t *testing.T) {
		repository, _ := createAdminUserRepositoryWithMockDB()
		assert.NotNil(t, repository)
	})
}

func TestAdminUserRepository_Delete(t *testing.T) {
	t.Run("Delete method on AdminUserRepository works properly", func(t *testing.T) {
		repository, mock := createAdminUserRepositoryWithMockDB()
		id, _ := uuid.NewUUID()
		mock.ExpectExec("DELETE FROM \"admin_users\" AS \"admin_users\"").WillReturnResult(sqlmock.NewResult(0, 0))
		err := repository.Delete(context.Background(), id)
		assert.NoError(t, err)
	})
}
