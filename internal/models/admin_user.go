package models

import (
	"github.com/bxcodec/faker/v4"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type AdminUser struct {
	bun.BaseModel `bun:"table:admin_users,alias:admin_users"`
	ID            uuid.UUID `bun:",pk,type:uuid,nullzero,notnull,default:uuid_generate_v4()"`
	CreatedAt     time.Time `bun:"type:timestamptz,nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:"type:timestamptz,nullzero,notnull,default:current_timestamp"`
	Name          string    `bun:",nullzero,notnull" faker:"name"`
	Email         string    `bun:",nullzero,notnull,unique" faker:"email"`
	Password      string    `bun:",nullzero,notnull"`
	Roles         []string  `bun:"type:jsonb,nullzero,notnull"`
}

// GetFakeAdminUser ... get fake AdminUser model
func GetFakeAdminUser() (*AdminUser, error) {
	entity := &AdminUser{}
	err := faker.FakeData(entity)

	if err != nil {
		return nil, err
	}

	entity.ID = uuid.New()
	entity.CreatedAt = time.Now()
	entity.UpdatedAt = time.Now()
	entity.Roles = []string{""}

	return entity, nil
}
