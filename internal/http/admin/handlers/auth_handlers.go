package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/takaaki-mizuno/go-gin-template/internal/http/admin/requests"
	adminResponses "github.com/takaaki-mizuno/go-gin-template/internal/http/admin/responses"
	commonResponses "github.com/takaaki-mizuno/go-gin-template/internal/http/common/responses"
	"net/http"
)

// AuthSigninPost ... endpoint for POST /auth/authorize
func (handler *Handler) AuthSigninPost(c *gin.Context) {
	var update requests.SigninPost
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusInternalServerError, commonResponses.NewInvalidParameterErrorStatus(err))
		return
	}
	accessToken, adminUser, err := handler.adminUserService.GetTokenFromEmailAndPassword(c.Request.Context(), update.Email, update.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, commonResponses.NewInternalServerError(err))
		return
	}

	response := adminResponses.NewToken(*accessToken, *adminUser)
	c.JSON(http.StatusOK, response)
}
