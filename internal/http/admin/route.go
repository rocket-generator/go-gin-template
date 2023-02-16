package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/takaaki-mizuno/go-gin-template/internal/http/admin/handlers"
	adminMiddlewares "github.com/takaaki-mizuno/go-gin-template/internal/http/admin/middlewares"
	"go.uber.org/dig"
)

func SetupRoutes(app *gin.Engine, container *dig.Container) error {
	var adminHandler handlers.HandlerInterface
	var adminAuthMiddleware gin.HandlerFunc

	if err := container.Invoke(func(
		h handlers.HandlerInterface,
		_adminMiddlewares adminMiddlewares.Middlewares,
	) {
		adminHandler = h
		adminAuthMiddleware = _adminMiddlewares.Auth
	}); err != nil {
		return err
	}

	adminRouter := app.Group("api/admin")
	adminRouter.POST("/auth/signin", adminHandler.AuthSigninPost)

	adminAuthRequiredRouter := adminRouter.Group("")
	adminAuthRequiredRouter.Use(adminAuthMiddleware)

	adminAuthRequiredRouter.GET("/admin_users", adminHandler.AdminUsersGet)
	adminAuthRequiredRouter.POST("/admin_users", adminHandler.AdminUsersPost)
	adminAuthRequiredRouter.GET("/admin_users/:id", adminHandler.AdminUserGet)
	adminAuthRequiredRouter.PUT("/admin_users/:id", adminHandler.AdminUserPut)
	adminAuthRequiredRouter.DELETE("/admin_users/:id", adminHandler.AdminUserDelete)

	return nil
}
