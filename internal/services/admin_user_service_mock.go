package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/takaaki-mizuno/go-gin-template/internal/models"
)

type AdminUserServiceMock struct{}

func (service *AdminUserServiceMock) GetTokenFromEmailAndPassword(ctx context.Context, email string, password string) (*string, *models.AdminUser, error) {
	model, _ := models.GetFakeAdminUser()
	token := ""
	return &token, model, nil
}

func (service *AdminUserServiceMock) GetAdminUserFromToken(ctx context.Context, token string) (*models.AdminUser, error) {
	model, _ := models.GetFakeAdminUser()
	return model, nil
}

func (service *AdminUserServiceMock) CreateAdminUser(ctx context.Context, name string, email string, password string) (*models.AdminUser, error) {
	model, _ := models.GetFakeAdminUser()
	model.Name = name
	model.Email = email
	return model, nil
}

func (service *AdminUserServiceMock) UpdateAdminUser(ctx context.Context, id uuid.UUID, name *string, email *string, password *string) (*models.AdminUser, error) {
	model, _ := models.GetFakeAdminUser()
	model.ID = id
	if name != nil {
		model.Name = *name
	}
	if email != nil {
		model.Email = *email
	}
	return model, nil
}

func (service *AdminUserServiceMock) DeleteAdminUser(ctx context.Context, id uuid.UUID) error {
	return nil
}
