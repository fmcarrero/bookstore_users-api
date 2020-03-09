package app

import (
	"github.com/fmcarrero/bookstore_users-api/infrastructure/controllers"
)

func mapUrls(handler controllers.RedirectUserHandler) {

	router.GET("/ping", controllers.Ping)

	router.POST("/users", handler.CreateUser)
	router.GET("/users/:user_id", handler.GetUser)
}
