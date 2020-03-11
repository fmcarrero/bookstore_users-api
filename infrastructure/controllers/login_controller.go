package controllers

import (
	"github.com/fmcarrero/bookstore_users-api/application/commands"
	"github.com/fmcarrero/bookstore_users-api/application/usescases"
	"github.com/fmcarrero/bookstore_users-api/infrastructure/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginUserHandler interface {
	Login(c *gin.Context)
}
type HandlerLogin struct {
	LoginUseCase usescases.LoginUseCase
}

func (h *HandlerLogin) Login(c *gin.Context) {
	var loginCommand commands.LoginCommand
	if err := c.ShouldBindJSON(&loginCommand); err != nil {
		restErr := errors.NewBadRequest("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, loginErr := h.LoginUseCase.Handler(loginCommand)
	if loginErr != nil {
		restErr := errors.NewBadRequest(loginErr.Error())
		c.JSON(http.StatusNotFound, restErr)
		return
	}

	c.JSON(http.StatusOK, result)
}
