package repository

import (
	"database/sql"
	"fmt"
	"log"
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
			log.Fatalf("error creating database directory: %v", err)
		}
	}

	DB, err := sql.Open(config.Driver, config.DbPath)

	if err != nil {
		log.Fatal(err)
	}

	createTable(DB, config.CreateUserTable)
	createTable(DB, config.CreateFollowsTable)
	createTable(DB, config.CreateTweetsTable)

	return Repository{db: DB}, nil
}

func createTable(db *sql.DB, query string) error {
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("error creating table. Error: %v", err)
	}
	return nil
}
