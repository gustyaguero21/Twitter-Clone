package repository

import (
	"database/sql"
	"fmt"
	"os"
	"twitter-clone/cmd/config"

	_ "modernc.org/sqlite"
)

type Repository struct {
	db *sql.DB
}

func NewRepository() (Repository, error) {

	dataDir := "./internal/data"
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		if err := os.Mkdir(dataDir, os.ModePerm); err != nil {
			return Repository{}, fmt.Errorf("error creating database directory. Error: %v", err)
		}
	}

	DB, err := sql.Open(config.Driver, config.DbPath)

	if err != nil {
		return Repository{}, fmt.Errorf("error opening  database. Error: %v", err)
	}

	if err := DB.Ping(); err != nil {
		return Repository{}, fmt.Errorf("error connecting with database. Error: %v", err)
	}

	tables := []string{config.CreateUserTable, config.CreateFollowerTable, config.CreateTweetsTable}

	for _, table := range tables {
		if err := createTable(DB, table); err != nil {
			return Repository{}, fmt.Errorf("error creating table. Error: %v", err)
		}
	}

	return Repository{db: DB}, nil
}

func createTable(db *sql.DB, query string) error {
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("error creating table. Error: %v", err)
	}
	return nil
}
