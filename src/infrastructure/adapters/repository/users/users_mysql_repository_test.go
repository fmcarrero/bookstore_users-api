package users

import (
	"context"
	"fmt"
	"github.com/fmcarrero/bookstore_users-api/src/domain/model"
	"github.com/fmcarrero/bookstore_users-api/src/infrastructure/adapters/repository/models"
	"github.com/fmcarrero/bookstore_users-api/src/infrastructure/database_client"
	"github.com/fmcarrero/bookstore_users-api/src/infrastructure/mappers/users_mapper"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"os"
	"testing"
)

var (
	userMysqlRepository UserMysqlRepository
)

func TestMain(m *testing.M) {
	fmt.Println("about to start oauth tests")
	containerMockServer, ctx := load()
	code := m.Run()
	beforeAll(containerMockServer, ctx)
	os.Exit(code)
}
func load() (testcontainers.Container, context.Context) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "mysql:latest",
		ExposedPorts: []string{"3306/tcp", "33060/tcp"},
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": "secret",
			"MYSQL_DATABASE":      "users_db",
		},
		WaitingFor: wait.ForLog("port: 3306  MySQL Community Server - GPL"),
	}
	mysqlC, _ := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	host, _ := mysqlC.Host(ctx)
	p, _ := mysqlC.MappedPort(ctx, "3306/tcp")
	port := p.Port()
	_ = os.Setenv("MYSQL_USERS_HOST", host)
	_ = os.Setenv("MYSQL_USERS_PORT", port)
	_ = os.Setenv("MYSQL_USERS_SCHEMA", "users_db")
	_ = os.Setenv("MYSQL_USERS_USERNAME", "root")
	_ = os.Setenv("MYSQL_USERS_PASSWORD", "secret")

	userMysqlRepository = UserMysqlRepository{
		Db: database_client.GetDatabaseInstance(),
	}
	return mysqlC, ctx
}
func beforeAll(container testcontainers.Container, ctx context.Context) {
	_ = container.Terminate(ctx)
}
func TestUserMysqlRepository_Save(t *testing.T) {
	tx := userMysqlRepository.Db.Begin()
	defer tx.Rollback()
	var user model.User
	user, _ = user.CreateUser("Franklin", "Carrero", "mauriciocarrero15@gmail.com", "sistemas31")
	err := userMysqlRepository.Save(&user)

	assert.Nil(t, err)
	assert.EqualValues(t, user.FirstName, "Franklin", "user names are differences")
	assert.NotEqual(t, user.Password, "sistemas31")
	assert.NotNil(t, user.Id, "user shouldn't be nil ")
}

func TestUserMysqlRepository_Get(t *testing.T) {

	tx := userMysqlRepository.Db.Begin()
	defer tx.Rollback()
	var user model.User
	user, _ = user.CreateUser("Franklin", "Carrero", "mauriciocarrero15@gmail.com", "sistemas31")
	var userDb models.UserDb
	userDb = users_mapper.UserToUserDb(user)
	if err := userMysqlRepository.Db.Create(&userDb).Error; err != nil {
		assert.Fail(t, err.Error())
	}
	user, err := userMysqlRepository.Get(userDb.ID)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, userDb.ID, user.Id)
	assert.EqualValues(t, "Carrero", user.LastName)
}
