package repository

import (
	"database/sql"
	"fmt"
	"twitter-clone/cmd/config"
	"twitter-clone/internal/models"
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

func (fr *FollowerRepository) SearchFollowers(username string) ([]models.Followers, error) {

	var following_list []models.Followers

	rows, err := fr.db.Query(config.ShowFollowersQuery, username)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var followings models.Followers

		if err := rows.Scan(&followings.FollowingUsername); err != nil {
			return nil, err
		}
		following_list = append(following_list, followings)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return following_list, nil
}
