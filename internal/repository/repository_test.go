package repository

import (
	"database/sql"
	"testing"
	"twitter-clone/cmd/config"

	"github.com/stretchr/testify/assert"
)

func Test_CreateTable_Success(t *testing.T) {
	//given

	db, err := sql.Open(config.Driver, ":memory:")

	if err != nil {
		t.Fatal("Error creating database in memory")
	}

	defer db.Close()

	//act
	err = createTable(db, config.CreateUserTable)

	//asserts
	assert.NoError(t, err)
}

func Test_CreateTable_Error(t *testing.T) {
	//given
	query := `CREATE error IF NOT EXISTS users (id PRIMARY KEY);`

	db, err := sql.Open(config.Driver, ":memory:")

	if err != nil {
		t.Fatal("Error creating database in memory")
	}

	defer db.Close()

	//act
	err = createTable(db, query)

	//asserts
	assert.Error(t, err)
}
