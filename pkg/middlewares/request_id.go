package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RequestID ... Add request ID on each request
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.New().String()
		c.Set("RequestID", requestID)
		c.Writer.Header().Set("X-Request-Id", requestID)
		c.Next()
	}
}
