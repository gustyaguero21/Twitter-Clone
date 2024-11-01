package repository

import (
	"database/sql"
	"fmt"
	"twitter-clone/cmd/config"
	"twitter-clone/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(repo Repository) *UserRepository {
	return &UserRepository{db: repo.db}
}

func (ur *UserRepository) SaveUser(user models.Users) error {
	_, err := ur.db.Exec(config.SaveUserQuery, user.ID, user.Username)

	if err != nil {
		return fmt.Errorf("error saving user on database. Error: %v", err)
	}
	return nil
}
