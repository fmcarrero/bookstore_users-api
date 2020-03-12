package controllers

import (
	"github.com/fmcarrero/bookstore_oauth-go/oauth"
	"github.com/fmcarrero/bookstore_users-api/application/commands"
	"github.com/fmcarrero/bookstore_users-api/application/usescases"
	"github.com/fmcarrero/bookstore_users-api/infrastructure/marshallers"
	"github.com/fmcarrero/bookstore_users-api/infrastructure/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RedirectUserHandler interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	FindByStatus(c *gin.Context)
}

type Handler struct {
	CreatesUseCase          usescases.CreatesUserPort
	GetUserUseCase          usescases.GetUserUseCase
	UseCaseUpdateUser       usescases.UpdateUserUseCase
	UseCaseDeleteUser       usescases.DeleteUserUseCase
	UseCaseFindUserByStatus usescases.FindUsersByStatusUseCase
}

func (h *Handler) Create(c *gin.Context) {

	var userCommand commands.UserCommand
	if err := c.ShouldBindJSON(&userCommand); err != nil {
		restErr := errors.NewBadRequest("invalid json")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, createUserErr := h.CreatesUseCase.Handler(userCommand)
	if createUserErr != nil {
		restErr := errors.NewBadRequest(createUserErr.Error())
		c.JSON(http.StatusBadRequest, restErr)
		return
	}
	isPublic := c.GetHeader("X-Public") == "true"
	c.JSON(http.StatusCreated, marshallers.Marshall(isPublic, result))
}
func (h *Handler) Get(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		restErr := errors.NewBadRequest("user_id should be valid")
		c.JSON(restErr.Status, restErr)
		return
	}
	user, errGet := h.GetUserUseCase.Handler(userId)
	if errGet != nil {
		restErr := errors.NewBadRequest(errGet.Error())
		c.JSON(restErr.Status, restErr)
		return
	}

	if oauth.GetCallerId(c.Request) == user.Id {
		c.JSON(http.StatusOK, marshallers.Marshall(false, user))
		return
	}
	c.JSON(http.StatusOK, marshallers.Marshall(oauth.IsPublic(c.Request), user))

}

func (h *Handler) Update(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		restErr := errors.NewBadRequest("user_id should be valid")
		c.JSON(restErr.Status, restErr)
		return
	}
	var userCommand commands.UserCommand
	if err := c.ShouldBindJSON(&userCommand); err != nil {
		restErr := errors.NewBadRequest("invalid json")
		c.JSON(restErr.Status, restErr)
		return
	}
	user, updateErr := h.UseCaseUpdateUser.Handler(userId, userCommand)
	if updateErr != nil {
		restErr := errors.NewBadRequest(updateErr.Error())
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, &user)
}

func (h *Handler) Delete(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		restErr := errors.NewBadRequest("user_id should be valid")
		c.JSON(restErr.Status, restErr)
		return
	}
	errDelete := h.UseCaseDeleteUser.Handler(userId)
	if errDelete != nil {
		restErr := errors.NewBadRequest(errDelete.Error())
		c.JSON(restErr.Status, restErr)
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *Handler) FindByStatus(c *gin.Context) {
	status := c.Query("status")
	users, err := h.UseCaseFindUserByStatus.Handler(status)
	if err != nil {
		restErr := errors.NewBadRequest(err.Error())
		c.JSON(restErr.Status, restErr)
		return
	}
	isPublic := c.GetHeader("X-Public") == "true"
	c.JSON(http.StatusOK, marshallers.MarshallArray(isPublic, users))
}
