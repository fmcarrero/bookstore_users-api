package authentication

import (
	"github.com/fmcarrero/bookstore_oauth-go/oauth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := oauth.AuthenticateRequest(c.Request); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			c.Abort()
			return
		}
		// add session verification here, like checking if the user and authType
		// combination actually exists if necessary. Try adding caching this (redis)
		// since this middleware might be called a lot
		c.Next()
	}
}
