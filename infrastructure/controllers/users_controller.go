package controllers

import (
	"fmt"
	"github.com/fmcarrero/bookstore_users-api/application/commands"
	"github.com/fmcarrero/bookstore_users-api/application/usescases"
	"github.com/fmcarrero/bookstore_users-api/infrastructure/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RedirectUserHandler interface {
	CreateUser(c *gin.Context)
	GetUser(c *gin.Context)
}

type Handler struct {
	CreatesUseCase usescases.CreatesUseCase
}

func (h *Handler) CreateUser(c *gin.Context) {

	var userCommand commands.UserCommand
	if err := c.ShouldBindJSON(&userCommand); err != nil {
		restErr := errors.NewBadRequest("invalid json")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, createUserErr := h.CreatesUseCase.Handler(userCommand)
	if createUserErr != nil {
		c.JSON(http.StatusBadRequest, createUserErr.Error())
		return
	}
	c.JSON(http.StatusCreated, result)
}
func (h *Handler) GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		restErr := errors.NewBadRequest("user_id should be valid")
		c.JSON(restErr.Status, restErr)
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("%v", userId))
}
