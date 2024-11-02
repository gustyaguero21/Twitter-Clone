package repository

import (
	"database/sql"
	"testing"
	"twitter-clone/cmd/config"

	"github.com/stretchr/testify/assert"
)

func TestCreateTable_Success(t *testing.T) {
	//given
	query := `CREATE TABLE IF NOT EXISTS users (id PRIMARY KEY);`

	db, err := sql.Open(config.Driver, ":memory:")

	if err != nil {
		t.Fatal("Error creating database in memory")
	}

	//act
	err = createTable(db, query)

	//asserts
	assert.NoError(t, err)
}

func TestCreateTable_Error(t *testing.T) {
	//given
	query := `CREATE error IF NOT EXISTS users (id PRIMARY KEY);`

	db, err := sql.Open(config.Driver, ":memory:")

	if err != nil {
		t.Fatal("Error creating database in memory")
	}

	//act
	err = createTable(db, query)

	//asserts
	assert.Error(t, err)
}
