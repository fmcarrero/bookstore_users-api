package app

import (
	"fmt"
	"github.com/fmcarrero/bookstore_users-api/application/usescases"
	"github.com/fmcarrero/bookstore_users-api/domain/ports"
	"github.com/fmcarrero/bookstore_users-api/infrastructure/adapters/repository"
	"github.com/fmcarrero/bookstore_users-api/infrastructure/adapters/repository/models"
	"github.com/fmcarrero/bookstore_users-api/infrastructure/controllers"
	"github.com/fmcarrero/bookstore_users-api/infrastructure/logger"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"os"
)

const (
	MysqlUsersUsername = "MYSQL_USERS_USERNAME"
	MysqlUsersPassword = "MYSQL_USERS_PASSWORD"
	MysqlUsersHost     = "MYSQL_USERS_HOST"
	MysqlUsersSchema   = "MYSQL_USERS_SCHEMA"
)

var (
	router = gin.Default()
)

func StartApplication() {
	var handler = createHandler()
	mapUrls(handler)
	logger.Info("about to start the application")
	_ = router.Run(":8080")
}
func createHandler() controllers.RedirectUserHandler {
	userRepository := getUsersRepository()
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
	return &repository.UserMysqlRepository{
		Db: getDatabaseInstance(),
	}
}

func getDatabaseInstance() *gorm.DB {
	_ = godotenv.Load()
	userName := os.Getenv(MysqlUsersUsername)
	password := os.Getenv(MysqlUsersPassword)
	host := os.Getenv(MysqlUsersHost)
	schema := os.Getenv(MysqlUsersSchema)
	dataSource := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=UTC", userName, password, host, schema)
	db, err := gorm.Open("mysql", dataSource)
	if err != nil {
		fmt.Println(err)
		db.Close()
		panic("database not working")
	}
	db.SingularTable(true)
	migrateDatabase(db)

	return db
}

func migrateDatabase(db *gorm.DB) {
	db.AutoMigrate(&models.UserDb{})
}
