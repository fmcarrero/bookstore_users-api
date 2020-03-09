package app

import (
	"fmt"
	"github.com/fmcarrero/bookstore_users-api/application/usescases"
	"github.com/fmcarrero/bookstore_users-api/domain/ports"
	"github.com/fmcarrero/bookstore_users-api/infrastructure/adapters/repository"
	"github.com/fmcarrero/bookstore_users-api/infrastructure/adapters/repository/models"
	"github.com/fmcarrero/bookstore_users-api/infrastructure/controllers"
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
	_ = router.Run(":8080")
}
func createHandler() controllers.RedirectUserHandler {
	return newHandler(newCreatesUseCase(getUsersRepository()))
}
func newCreatesUseCase(repository ports.UsersRepository) usescases.CreatesUseCase {
	return &usescases.UseCaseUserCreate{
		UserRepository: repository,
	}
}
func newHandler(createUser usescases.CreatesUseCase) controllers.RedirectUserHandler {
	return &controllers.Handler{CreatesUseCase: createUser}
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
