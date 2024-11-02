package repository

import (
	"database/sql"
	"fmt"
	"twitter-clone/cmd/config"
	"twitter-clone/internal/models"
)

type FollowerRepository struct {
	Db *sql.DB
}

func NewFollowerRepository(repo Repository) *FollowerRepository {
	return &FollowerRepository{Db: repo.Db}
}

func (fr *FollowerRepository) Follow(follower_username, following_username string) error {

	_, err := fr.Db.Exec(config.FollowUserQuery, follower_username, following_username)

	if err != nil {
		return fmt.Errorf("user already follow this user. Error: %v", err)
	}
	return nil
}

func (fr *FollowerRepository) SearchFollowers(username string) ([]models.Followers, error) {

	var following_list []models.Followers

	rows, err := fr.Db.Query(config.ShowFollowersQuery, username)

	if err != nil {
		return nil, fmt.Errorf("error searching followers. Error: %v", err)
	}

	defer rows.Close()

	for rows.Next() {

		var followings models.Followers

		if err := rows.Scan(&followings.FollowingUsername); err != nil {
			return nil, fmt.Errorf("error scanning rows. Error: %v", err)
		}
		following_list = append(following_list, followings)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return following_list, nil
}
