package handlers

import (
	"github.com/takaaki-mizuno/go-gin-template/config"
	"github.com/uptrace/bun"
	"go.uber.org/zap"
)

// HandlerInterface ...
type HandlerInterface interface {
	// {{ START HANDLER DEPENDENCY }}
	// Template: {{ .Name.Plural.Title }}GetHandler(ctx *gin.Context)
	// Template: {{ .Name.Singular.Title }}GetHandler(ctx *gin.Context)
	// Template: {{ .Name.Plural.Title }}PostHandler(ctx *gin.Context)
	// Template: {{ .Name.Singular.Title }}PutHandler(ctx *gin.Context)
	// Template: {{ .Name.Singular.Title }}DeleteHandler(ctx *gin.Context)
	// {{ END HANDLER DEPENDENCY}}
}

// Handler ...
type Handler struct {
	db     *bun.DB
	config *config.Config
	logger *zap.Logger
	// {{ START HANDLER DEPENDENCY }}
	// Template: {{ .Name.Singular.Title }}DeleteHandler(ctx *gin.Context)
	// {{ END HANDLER DEPENDENCY}}
}

// NewHandler ...
func NewHandler(
	db *bun.DB,
	config *config.Config,
	logger *zap.Logger,
	// {{ START HANDLER DEPENDENCY ARGUMENTS }}
	// {{ END HANDLER DEPENDENCY ARGUMENTS }}
) HandlerInterface {
	return &Handler{
		db:     db,
		config: config,
		logger: logger,
		// {{ START HANDLER DEPENDENCY ASSIGNMENT }}
		// {{ START HANDLER DEPENDENCY ASSIGNMENT }}
	}
}
