package repository

import (
	"database/sql"
	"testing"
	"twitter-clone/cmd/config"

	"github.com/stretchr/testify/assert"
)

func Test_Follow_Success(t *testing.T) {

	//given

	db, err := sql.Open(config.Driver, ":memory:")

	if err != nil {
		t.Fatal("Error creating database in memory")
	}

	defer db.Close()

	_, err = db.Exec(config.CreateFollowerTable)

	if err != nil {
		t.Fatal("Error creating followers table")
	}

	followRepo := NewFollowerRepository(Repository{Db: db})

	//act

	err = followRepo.Follow("pepe", "alicia")

	//asserts

	assert.NoError(t, err)
}

func Test_Follow_Error(t *testing.T) {

	//given

	db, err := sql.Open(config.Driver, ":memory:")

	if err != nil {
		t.Fatal("Error creating database in memory")
	}

	defer db.Close()

	_, err = db.Exec(config.CreateFollowerTable)

	if err != nil {
		t.Fatal("Error creating followers table")
	}

	followRepo := NewFollowerRepository(Repository{Db: db})

	//act

	err = followRepo.Follow("pepe", "alicia")

	if err != nil {
		t.Fatal("Error following user")
	}

	err = followRepo.Follow("pepe", "alicia")
	//asserts

	assert.Error(t, err)
}
