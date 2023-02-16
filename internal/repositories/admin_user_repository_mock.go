package repositories

import (
	"context"
	"github.com/google/uuid"
	"github.com/takaaki-mizuno/go-gin-template/internal/models"
)

// MockAdminUserRepository ...
type MockAdminUserRepository struct {
}

func (repository *MockAdminUserRepository) FindByID(ctx context.Context, id uuid.UUID) (*models.AdminUser, error) {
	model, _ := models.GetFakeAdminUser()
	model.ID = id
	return model, nil
}

func (repository *MockAdminUserRepository) FindByEmail(ctx context.Context, email string) (*models.AdminUser, error) {
	model, _ := models.GetFakeAdminUser()
	model.Email = email
	return model, nil
}

func (repository *MockAdminUserRepository) Get(ctx context.Context, filters *[]Filter, order string, direction string, offset int64, limit int64) (*[]models.AdminUser, int64, error) {
	model01, _ := models.GetFakeAdminUser()
	model02, _ := models.GetFakeAdminUser()
	return &[]models.AdminUser{*model01, *model02}, 2, nil
}

func (repository *MockAdminUserRepository) Save(ctx context.Context, model models.AdminUser) (*models.AdminUser, error) {
	return &model, nil
}

func (repository *MockAdminUserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}
