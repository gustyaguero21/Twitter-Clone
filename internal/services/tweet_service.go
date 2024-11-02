package services

import (
	"context"
	"fmt"
	"time"
	"twitter-clone/cmd/config"
	"twitter-clone/internal/models"
	"twitter-clone/internal/repository"

	"github.com/google/uuid"
)

type TweetServices struct {
	Repo *repository.Repository
}

func NewTweetService(repo *repository.Repository) *TweetServices {
	return &TweetServices{Repo: repo}
}

func (ts *TweetServices) CreatePost(ctx context.Context, user string, content string) (models.Tweets, error) {

	if err := ts.validate(content); err != nil {
		return models.Tweets{}, err
	}
	newPost := models.Tweets{
		ID:       uuid.New(),
		User:     user,
		Content:  content,
		PostedAt: time.Now(),
	}
	tweetRepo := repository.NewPostRepository(*ts.Repo)

	if err := tweetRepo.SavePost(newPost.ID, newPost.User, newPost.Content, newPost.PostedAt); err != nil {
		return models.Tweets{}, err
	}

	return newPost, nil
}

func (ts *TweetServices) validate(content string) error {

	if len(content) > config.MaxContentLength {
		return fmt.Errorf("content cannot exceed 280 characters")
	}
	return nil
}

func (ts *TweetServices) ShowTimeline(username string) ([]models.Tweets, error) {

	tweetRepo := repository.NewPostRepository(*ts.Repo)

	tweetList, tweetErr := tweetRepo.GetTimeline(username)

	if tweetErr != nil {
		return []models.Tweets{}, tweetErr
	}
	return tweetList, nil

}
