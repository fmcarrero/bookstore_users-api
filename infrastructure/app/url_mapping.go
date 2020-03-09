package app

import (
	"github.com/fmcarrero/bookstore_users-api/infrastructure/controllers"
)

func mapUrls(handler controllers.RedirectUserHandler) {

	router.GET("/ping", controllers.Ping)

	router.POST("/users", handler.Create)
	router.GET("/users/:user_id", handler.Get)
	router.PUT("/users/:user_id", handler.Update)
	router.DELETE("/users/:user_id", handler.Delete)
	router.GET("/internal/users/search", handler.FindByStatus)
}
