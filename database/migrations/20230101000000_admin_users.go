package migrations

import (
	"context"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type AdminUser struct {
	ID        uuid.UUID `bun:",pk,type:uuid,nullzero,notnull,default:uuid_generate_v4()"`
	CreatedAt time.Time `bun:"type:timestamptz,nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"type:timestamptz,nullzero,notnull,default:current_timestamp"`
	Name      string    `bun:",nullzero,notnull"`
	Email     string    `bun:",nullzero,notnull,unique"`
	Password  string    `bun:",nullzero,notnull"`
	Roles     []string  `bun:"type:jsonb,nullzero,notnull"`
}

func init() {
	registerCreateModelMigration(&AdminUser{})
}

func registerCreateModelMigration(model interface{}) {

	up := func(ctx context.Context, db *bun.DB) error {
		if _, err := db.
			NewCreateTable().
			Model(model).
			IfNotExists().
			WithForeignKeys().
			Exec(ctx); err != nil {
			return err
		}

		return nil
	}

	down := func(ctx context.Context, db *bun.DB) error {
		if _, err := db.
			NewDropTable().
			Model(model).
			IfExists().
			Exec(ctx); err != nil {
			return err
		}

		return nil
	}

	Migrations.MustRegister(up, down)
}
