package middlewares

import (
	"github.com/gin-gonic/gin"
	commonResponse "github.com/takaaki-mizuno/go-gin-template/internal/http/common/responses"
	"github.com/takaaki-mizuno/go-gin-template/internal/services"
	"github.com/uptrace/bun"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

// Auth ... Get Auth Middleware
func Auth(
	adminUserService services.AdminUserServiceInterface,
	logger *zap.Logger,
	database *bun.DB,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			c.JSON(http.StatusUnauthorized, commonResponse.NewUnauthorizedError("No Authorization header provided"))
			c.Abort()
			return
		}
		token := strings.TrimPrefix(auth, "Bearer ")
		if token == auth {
			c.JSON(http.StatusUnauthorized, commonResponse.NewUnauthorizedError("Could not find bearer token in Authorization header"))
			c.Abort()
			return
		}
		user, err := adminUserService.GetAdminUserFromToken(c, token)
		if err != nil {
			c.JSON(http.StatusInternalServerError, commonResponse.NewInternalServerError(err))
			c.Abort()
			return
		}
		if user == nil {
			c.JSON(http.StatusUnauthorized, commonResponse.NewUnauthorizedError("Invalid token"))
			c.Abort()
			return
		}
		c.Set("adminUser", user)
		c.Next()
	}
}
