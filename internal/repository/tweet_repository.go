package repository

import (
	"database/sql"
	"fmt"
	"time"
	"twitter-clone/cmd/config"
	"twitter-clone/internal/models"

	"github.com/google/uuid"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(repo Repository) *PostRepository {
	return &PostRepository{db: repo.db}
}

func (pr *PostRepository) SavePost(id uuid.UUID, user string, content string, posted_at time.Time) error {

	_, err := pr.db.Exec(config.CreatePostQuery, id.String(), user, content, posted_at.Format("2006-01-02 15:04:05"))

	if err != nil {
		return fmt.Errorf("error creating post. Error: %v", err)
	}
	return nil
}

func (pr *PostRepository) GetTimeline(username string) ([]models.Tweets, error) {

	rows, err := pr.db.Query(config.TimelineQuery, username)

	if err != nil {
		return nil, fmt.Errorf("error fetching timeline tweets: %v", err)
	}
	defer rows.Close()

	var tweets []models.Tweets
	for rows.Next() {
		var tweet models.Tweets
		if err := rows.Scan(&tweet.User, &tweet.Content, &tweet.PostedAt); err != nil {
			return nil, err
		}
		tweets = append(tweets, tweet)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tweets, nil
}
