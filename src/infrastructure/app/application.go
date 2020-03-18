package app

import (
	"github.com/fmcarrero/bookstore_users-api/src/application/usescases"
	"github.com/fmcarrero/bookstore_users-api/src/domain/ports"
	"github.com/fmcarrero/bookstore_users-api/src/infrastructure/adapters/repository/users"
	"github.com/fmcarrero/bookstore_users-api/src/infrastructure/app/middlewares/error_handler"
	"github.com/fmcarrero/bookstore_users-api/src/infrastructure/controllers"
	"github.com/fmcarrero/bookstore_users-api/src/infrastructure/database_client"
	"github.com/fmcarrero/bookstore_utils-go/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	router = gin.Default()
)

func StartApplication() {

	_ = godotenv.Load()
	router.Use(error_handler.ErrorHandler())
	userRepository := getUsersRepository()
	var handler = createHandler(userRepository)
	mapUrls(handler)
	mapUrlLogin(createLoginHandler(userRepository))

	logger.Info("about to start the application")
	_ = router.Run(":8081")
}
func createLoginHandler(userRepository ports.UsersRepository) controllers.LoginUserHandler {
	useCaseLogin := &usescases.UseCaseLogin{UserRepository: userRepository}
	return &controllers.HandlerLogin{LoginUseCase: useCaseLogin}
}
func createHandler(userRepository ports.UsersRepository) controllers.RedirectUserHandler {

	return newHandler(newCreatesUseCase(userRepository), newGetUserUseCase(userRepository),
		newUpdateUserUseCase(userRepository), newDeleteUserUseCase(userRepository),
		newFindUsersByStatusUseCase(userRepository))
}
func newCreatesUseCase(repository ports.UsersRepository) usescases.CreatesUserPort {
	return &usescases.UseCaseUserCreate{
		UserRepository: repository,
	}
}

func newGetUserUseCase(repository ports.UsersRepository) usescases.GetUserUseCase {
	return &usescases.UseCaseGetUser{
		UserRepository: repository,
	}
}

func newUpdateUserUseCase(repository ports.UsersRepository) usescases.UpdateUserUseCase {
	return &usescases.UseCaseUpdateUser{
		UserRepository: repository,
	}
}

func newDeleteUserUseCase(usersRepository ports.UsersRepository) usescases.DeleteUserUseCase {
	return &usescases.UseCaseDeleteUser{
		UserRepository: usersRepository,
	}
}

func newFindUsersByStatusUseCase(usersRepository ports.UsersRepository) usescases.FindUsersByStatusUseCase {
	return &usescases.UseCaseFindUserByStatus{
		UserRepository: usersRepository,
	}
}

func newHandler(createUser usescases.CreatesUserPort, getUserUseCase usescases.GetUserUseCase, updateUserUseCase usescases.UpdateUserUseCase,
	deleteUserUseCase usescases.DeleteUserUseCase, useCaseFindUserByStatus usescases.FindUsersByStatusUseCase) controllers.RedirectUserHandler {
	return &controllers.Handler{CreatesUseCase: createUser, GetUserUseCase: getUserUseCase, UseCaseUpdateUser: updateUserUseCase,
		UseCaseDeleteUser:       deleteUserUseCase,
		UseCaseFindUserByStatus: useCaseFindUserByStatus,
	}
}
func getUsersRepository() ports.UsersRepository {
	return &users.UserMysqlRepository{
		Db: database_client.GetDatabaseInstance(),
	}
}
