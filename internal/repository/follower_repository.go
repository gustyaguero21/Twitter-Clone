package repository

import (
	"database/sql"
	"fmt"
	"twitter-clone/cmd/config"
)

type FollowerRepository struct {
	db *sql.DB
}

func NewFollowerRepository(repo Repository) *FollowerRepository {
	return &FollowerRepository{db: repo.db}
}

func (fr *FollowerRepository) Follow(follower_username, following_username string) error {
	_, err := fr.db.Exec(config.FollowUserQuery, follower_username, following_username)
	if err != nil {
		return fmt.Errorf("error: You already follow this user. Error: %v", err.Error())
	}
	return nil
}
