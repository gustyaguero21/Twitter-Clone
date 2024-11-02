package repository

import (
	"database/sql"
	"fmt"
	"time"
	"twitter-clone/cmd/config"

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
