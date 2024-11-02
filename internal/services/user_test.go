package services

import (
	"context"
	"database/sql"
	"testing"
	"twitter-clone/cmd/config"
	"twitter-clone/internal/models"
	"twitter-clone/internal/repository"

	"github.com/stretchr/testify/assert"
)

func Test_CreateUser_Success(t *testing.T) {
	//given
	ctx := context.Background()
	db, err := sql.Open(config.Driver, ":memory:")

	if err != nil {
		t.Fatal("Error creating database in memory")
	}

	defer db.Close()

	_, err = db.Exec(config.CreateUserTable)

	if err != nil {
		t.Fatal("Error creating user table")
	}

	userService := NewUserService(&repository.Repository{Db: db})

	user := models.Users{
		Username: "pepe",
	}

	//act

	createdUser, err := userService.CreateUser(ctx, user)

	//asserts

	assert.NoError(t, err)
	assert.Equal(t, user.Username, createdUser.Username)
}
