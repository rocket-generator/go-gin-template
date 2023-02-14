package handlers

import (
	"github.com/takaaki-mizuno/go-gin-template/config"
	"github.com/uptrace/bun"
	"go.uber.org/zap"
)

// HandlerInterface ...
type HandlerInterface interface {
	// {{ START HANDLER DEPENDENCY }}
	// {{ END HANDLER DEPENDENCY}}
}

// Handler ...
type Handler struct {
	// {{ START SERVICE DEPENDENCY }}
	// Template: {{ .Name.Singular.Camel }}Service    services.{{ .Name.Singular.Title }}Service
	// {{ END SERVICE DEPENDENCY}}
	db     *bun.DB
	config *config.Config
	logger *zap.Logger
}

// NewHandler ...
func NewHandler(
	db *bun.DB,
	config *config.Config,
	logger *zap.Logger,
) HandlerInterface {
	return &Handler{
		db:     db,
		config: config,
		logger: logger,
	}
}
