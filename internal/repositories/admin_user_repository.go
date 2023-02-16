package repositories

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/takaaki-mizuno/go-gin-template/internal/models"
	"github.com/uptrace/bun"
	"go.uber.org/zap"
)

// AdminUserRepositoryInterface ...
type AdminUserRepositoryInterface interface {
	FindByID(ctx context.Context, id uuid.UUID) (*models.AdminUser, error)
	FindByEmail(ctx context.Context, email string) (*models.AdminUser, error)
	Get(ctx context.Context, filters *[]Filter, order string, direction string, offset int64, limit int64) (*[]models.AdminUser, int64, error)
	Save(ctx context.Context, model models.AdminUser) (*models.AdminUser, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

// AdminUserRepository ...
type AdminUserRepository struct {
	logger   *zap.Logger
	database *bun.DB
}

// FindByID ... Get by User
func (repository *AdminUserRepository) FindByID(ctx context.Context, id uuid.UUID) (*models.AdminUser, error) {
	var model models.AdminUser

	q := repository.database.
		NewSelect().
		Model(&model).
		Where("? = ?", bun.Ident("admin_users.id"), id)
	err := q.Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &model, nil
}

// FindByEmail ... Get by User
func (repository *AdminUserRepository) FindByEmail(ctx context.Context, email string) (*models.AdminUser, error) {
	var model models.AdminUser

	q := repository.database.
		NewSelect().
		Model(&model).
		Where("? = ?", bun.Ident("admin_users.email"), email)
	err := q.Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &model, nil
}

// Get ... Get AdminUser List
func (repository *AdminUserRepository) Get(ctx context.Context, filters *[]Filter, order string, direction string, offset int64, limit int64) (*[]models.AdminUser, int64, error) {
	_models := &[]models.AdminUser{}
	if direction != "desc" {
		direction = "asc"
	}
	q := repository.database.
		NewSelect().
		Model(_models)
	if filters != nil {
		q = ApplyFilters("admin_users", q, *filters)
	}

	q = q.Limit(int(limit)).Offset(int(offset)).Order(order + " " + direction)
	count, err := q.ScanAndCount(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, 0, nil
		}
		return nil, 0, err
	}

	return _models, int64(count), nil
}

// Save ... Save AdminUser
func (repository *AdminUserRepository) Save(ctx context.Context, model models.AdminUser) (*models.AdminUser, error) {
	if model.ID == uuid.Nil {
		_, err := repository.database.NewInsert().Model(&model).Returning("*").Exec(ctx)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := repository.database.NewUpdate().Model(&model).
			Where("? = ?", bun.Ident("admin_users.id"), model.ID).
			Returning("*").Exec(ctx)
		if err != nil {
			return nil, err
		}
	}
	return &model, nil
}

// Delete ... Delete AdminUser
func (repository *AdminUserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := repository.database.NewDelete().Model((*models.AdminUser)(nil)).
		Where("? = ?", bun.Ident("admin_users.id"), id).Exec(ctx)

	if err != nil {
		return err
	}
	return nil
}

// NewAdminUserRepository ... Create new repository
func NewAdminUserRepository(logger *zap.Logger, database *bun.DB) AdminUserRepositoryInterface {
	return &AdminUserRepository{logger: logger, database: database}
}
