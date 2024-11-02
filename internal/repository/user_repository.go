package repository

import (
	"context"
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

	ctx := context.Background()

	if ur.existsUser(ctx, user.ID.String(), user.Username) {
		return fmt.Errorf("user with ID %s and username %s already exists", user.ID.String(), user.Username)
	}

	_, err := ur.db.Exec(config.SaveUserQuery, user.ID, user.Username)
	if err != nil {
		return fmt.Errorf("error saving user on database. Error: %v", err)
	}
	return nil
}

func (ur *UserRepository) existsUser(ctx context.Context, id, username string) bool {

	var count int

	err := ur.db.QueryRowContext(ctx, config.ExistUserQuery, id, username).Scan(&count)
	if err != nil {
		return false
	}

	return count > 0
}
