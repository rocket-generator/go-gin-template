package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/takaaki-mizuno/go-gin-template/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"go.uber.org/zap"
	"time"
)

// InitDatabase ... initializes the database connection.
func InitDatabase(configInstance *Config, logger *zap.Logger) (*bun.DB, error) {
	dbConfig := configInstance.Database

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbConfig.Postgres.User, dbConfig.Postgres.Password, dbConfig.Postgres.Host, dbConfig.Postgres.Port, dbConfig.Postgres.Name)
	sqldb := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithDSN(dsn),
		pgdriver.WithTimeout(dbConfig.Timeout*time.Second),
		pgdriver.WithDialTimeout(dbConfig.DialTimeout*time.Second),
		pgdriver.WithReadTimeout(dbConfig.ReadTimeout*time.Second),
		pgdriver.WithWriteTimeout(dbConfig.WriteTimeout*time.Second),
	))

	db := bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(NewLoggerQueryHook(logger))
	return db, nil
}

// PingDB ...
func PingDB(context context.Context, db *bun.DB, logger *zap.Logger) bool {
	var num int
	if err := db.QueryRowContext(context, "SELECT 1").Scan(&num); err != nil {
		logger.Error("Database error", zap.Error(err))
		return false
	}
	return true
}

// GetMockDB ... Get mock db connection for unit testing purpose
func GetMockDB() (*bun.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gdb := bun.NewDB(db, pgdialect.New())
	if err != nil {
		return nil, nil, err
	}

	return gdb, mock, nil
}
