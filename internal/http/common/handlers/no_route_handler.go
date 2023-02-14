package handlers

import (
	"github.com/gin-gonic/gin"
	commonResponses "github.com/takaaki-mizuno/go-gin-template/internal/http/common/responses"
	"net/http"
)

// NoRoute ... endpoint for 404
func (handler *Handler) NoRoute(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, commonResponses.NewNotFoundError("Endpoint not found"))
	return
}
