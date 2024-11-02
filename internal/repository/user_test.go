package repository

import (
	"database/sql"
	"testing"
	"twitter-clone/cmd/config"
	"twitter-clone/internal/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_SaveUser_Success(t *testing.T) {
	//given

	db, err := sql.Open(config.Driver, ":memory:")

	if err != nil {
		t.Fatal("Error creating database in memory")
	}

	defer db.Close()

	_, err = db.Exec(config.CreateUserTable)

	if err != nil {
		t.Fatal("Error creating user table")
	}

	userRepo := NewUserRepository(Repository{db: db})

	user := models.Users{
		ID:       uuid.New(),
		Username: "pepe",
	}
	//act

	err = userRepo.SaveUser(user)

	//asserts

	assert.Nil(t, err)

}

func Test_SaveUser_Error(t *testing.T) {
	//given

	db, err := sql.Open(config.Driver, ":memory:")

	if err != nil {
		t.Fatal("Error creating database in memory")
	}

	defer db.Close()

	_, err = db.Exec(config.CreateUserTable)

	if err != nil {
		t.Fatal("Error creating user table")
	}

	userRepo := NewUserRepository(Repository{db: db})

	_, err = db.Exec(config.SaveUserQuery, "1", "pepe")

	if err != nil {
		t.Fatal("Error saving user")
	}

	user := models.Users{
		ID:       uuid.New(),
		Username: "pepe",
	}
	//act

	err = userRepo.SaveUser(user)

	//asserts

	assert.Error(t, err)
	assert.Equal(t, "user already exists", err.Error())

}
