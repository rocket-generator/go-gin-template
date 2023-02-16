package handlers

import (
	"github.com/google/uuid"
	"github.com/takaaki-mizuno/go-gin-template/internal/http/admin/requests"
	adminResponse "github.com/takaaki-mizuno/go-gin-template/internal/http/admin/responses"
	commonResponse "github.com/takaaki-mizuno/go-gin-template/internal/http/common/responses"
	"go.uber.org/zap"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"net/http"
)

// AdminUsersGet ... endpoint for GET /admin_users
func (handler *Handler) AdminUsersGet(c *gin.Context) {
	// http://my.api.url/posts?_sort=title&_order=ASC&_start=0&_end=24&title=bar
	offset, err := strconv.Atoi(c.DefaultQuery("_start", "0"))
	if err != nil || offset < 0 {
		offset = 0
	}
	end, err := strconv.Atoi(c.DefaultQuery("_end", "20"))
	if err != nil || end < 0 {
		end = 20
	}
	limit := end - offset
	if err != nil || limit < 1 || limit > 100 {
		limit = 20
	}
	order := strings.ToLower(c.DefaultQuery("_sort", "id"))
	direction := strings.ToLower(c.DefaultQuery("_order", "ASC"))
	if err != nil {
		handler.logger.Error("Error parsing sort", zap.Error(err))
		order = "id"
		direction = "asc"
	}
	found := false
	for _, column := range []string{
		"id",
		"email",
	} {
		if column == order {
			found = true
		}
	}
	if !found {
		order = "id"
	}

	data, count, err := handler.adminUserRepository.Get(
		c,
		nil,
		order,
		direction,
		int64(offset),
		int64(limit),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, commonResponse.NewInternalServerError(err))
		return
	}

	c.Writer.Header().Set("X-Total-Count", strconv.FormatInt(count, 10))

	response := adminResponse.NewAdminUsers(data)
	c.JSON(http.StatusOK, response)
}

// AdminUserGet ... endpoint for GET /admin_users/{id}
func (handler *Handler) AdminUserGet(c *gin.Context) {

	idString := c.Param("id")
	id, err := uuid.Parse(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, commonResponse.NewInternalServerError(err))
		return
	}
	model, err := handler.adminUserRepository.FindByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, commonResponse.NewInternalServerError(err))
		return
	}
	response := adminResponse.NewAdminUser(*model)
	c.JSON(http.StatusOK, response)
}

// AdminUsersPost ... endpoint for POST /admin_users
func (handler *Handler) AdminUsersPost(c *gin.Context) {
	var request requests.AdminUserCreate
	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, commonResponse.NewInvalidParameterErrorStatus(err))
		return
	}
	resultModel, err := handler.adminUserService.CreateAdminUser(c, request.Name, request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, commonResponse.NewInternalServerError(err))
		return
	}
	response := adminResponse.NewAdminUser(*resultModel)
	c.JSON(http.StatusCreated, response)
}

// AdminUserPut ... endpoint for PUT /admin_users/{id}
func (handler *Handler) AdminUserPut(c *gin.Context) {
	idString := c.Param("id")
	id, err := uuid.Parse(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, commonResponse.NewInternalServerError(err))
		return
	}
	_, err = handler.adminUserRepository.FindByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, commonResponse.NewInternalServerError(err))
		return
	}
	var update requests.AdminUserUpdate
	if err := c.Bind(&update); err != nil {
		c.JSON(http.StatusBadRequest, commonResponse.NewInvalidParameterErrorStatus(err))
		return
	}

	resultModel, err := handler.adminUserService.UpdateAdminUser(c, id, update.Name, update.Email, update.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, commonResponse.NewInternalServerError(err))
		return
	}
	response := adminResponse.NewAdminUser(*resultModel)
	c.JSON(http.StatusCreated, response)
}

// AdminUserDelete ... endpoint for DELETE /admin_users/{id}
func (handler *Handler) AdminUserDelete(c *gin.Context) {
	idString := c.Param("id")
	id, err := uuid.Parse(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, commonResponse.NewInternalServerError(err))
		return
	}
	_, err = handler.adminUserRepository.FindByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, commonResponse.NewInternalServerError(err))
		return
	}

	err = handler.adminUserService.DeleteAdminUser(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, commonResponse.NewInternalServerError(err))
		return
	}
	response := commonResponse.NewSuccessStatus()
	c.JSON(http.StatusOK, response)
}
