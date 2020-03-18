package controllers

import (
	"github.com/fmcarrero/bookstore_users-api/src/application/commands"
	"github.com/fmcarrero/bookstore_users-api/src/application/usescases"
	"github.com/fmcarrero/bookstore_utils-go/rest_errors"
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
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}
	result, loginErr := h.LoginUseCase.Handler(loginCommand)
	if loginErr != nil {
		restErr := rest_errors.NewBadRequestError(loginErr.Error())
		c.JSON(restErr.Status(), restErr)
		return
	}

	c.JSON(http.StatusOK, result)
}
