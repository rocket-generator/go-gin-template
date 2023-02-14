package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/takaaki-mizuno/go-gin-template/config"
	"github.com/uptrace/bun"
	"go.uber.org/zap"
)

// HandlerInterface ...
type HandlerInterface interface {
	Healthz(ctx *gin.Context)
	IndexGet(ctx *gin.Context)
	NoRoute(ctx *gin.Context)
	NoMethod(ctx *gin.Context)
}

// Handler ...
type Handler struct {
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
