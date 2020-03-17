package app

import (
	"github.com/fmcarrero/bookstore_users-api/infrastructure/app/middlewares/authentication"
	"github.com/fmcarrero/bookstore_users-api/infrastructure/controllers"
)

func mapUrls(handler controllers.RedirectUserHandler) {

	router.GET("/ping", controllers.Ping)

	router.POST("/users", handler.Create)

	router.Use(authentication.AuthRequired())
	{
		router.GET("/users/:user_id", handler.Get)
	}

	router.PUT("/users/:user_id", handler.Update)
	router.DELETE("/users/:user_id", handler.Delete)
	router.GET("/internal/users/search", handler.FindByStatus)
}

func mapUrlLogin(handler controllers.LoginUserHandler) {
	router.POST("/users/login", handler.Login)
}
