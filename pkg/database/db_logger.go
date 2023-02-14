package database

import (
	"context"
	"database/sql"
	"github.com/uptrace/bun"
	"go.uber.org/zap"
)

// LoggerQueryHook implements logger that correspond to `bun.QueryHook` interface.
type LoggerQueryHook struct {
	logger *zap.Logger
}

// NewLoggerQueryHook initializes a new `LoggerQueryHook`.
func NewLoggerQueryHook(logger *zap.Logger) *LoggerQueryHook {
	return &LoggerQueryHook{
		logger: logger,
	}
}

// BeforeQuery provides a hook to correspond to `bun.QueryHook` interface.
func (h *LoggerQueryHook) BeforeQuery(ctx context.Context, event *bun.QueryEvent) context.Context {
	return ctx
}

// AfterQuery provides a hook to correspond to `bun.QueryHook` interface.
func (h *LoggerQueryHook) AfterQuery(ctx context.Context, event *bun.QueryEvent) {
	h.logger.Info(event.Query)

	switch event.Err {
	case nil, sql.ErrNoRows, sql.ErrTxDone:
		return
	}

	if event.Err != nil {
		h.logger.Error(event.Err.Error())
	}
}
