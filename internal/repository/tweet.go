package repository

import (
	"database/sql"
	"fmt"
	"twitter-clone/cmd/config"
	"twitter-clone/internal/models"
)

type PostRepository struct {
	Db *sql.DB
}

func NewPostRepository(repo Repository) *PostRepository {
	return &PostRepository{Db: repo.Db}
}

func (pr *PostRepository) SavePost(tweet models.Tweets) error {

	if tweet.Content == "" {
		return fmt.Errorf("content cannot be empty")
	}

	_, err := pr.Db.Exec(config.CreatePostQuery, tweet.ID, tweet.User, tweet.Content, tweet.PostedAt.Format("2006-01-02 15:04:05"))

	if err != nil {
		return fmt.Errorf("error creating post. Error: %v", err)
	}

	return nil
}

func (pr *PostRepository) GetTimeline(username string) ([]models.Tweets, error) {

	rows, err := pr.Db.Query(config.TimelineQuery, username)

	if err != nil {
		return nil, fmt.Errorf("error fetching timeline tweets. Error: %v", err)
	}
	defer rows.Close()

	var tweets []models.Tweets

	for rows.Next() {

		var tweet models.Tweets

		if err := rows.Scan(&tweet.ID, &tweet.User, &tweet.Content, &tweet.PostedAt); err != nil {
			return nil, fmt.Errorf("error searching tweets. Error: %v", err)
		}

		tweets = append(tweets, tweet)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tweets, nil
}
