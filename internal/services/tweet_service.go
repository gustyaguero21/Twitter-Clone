package services

import (
	"context"
	"fmt"
	"sort"
	"sync"
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

	followingUsernames, err := tweetRepo.GetFollowingUsernames(username)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	tweetsChan := make(chan []models.Tweets, len(followingUsernames))

	for _, followedUser := range followingUsernames {
		wg.Add(1)
		go func(user string) {
			defer wg.Done()
			tweets, err := tweetRepo.GetTimeline(user)
			if err == nil {
				tweetsChan <- tweets
			}
		}(followedUser)
	}

	go func() {
		wg.Wait()
		close(tweetsChan)
	}()

	var allTweets []models.Tweets

	for tweets := range tweetsChan {
		allTweets = append(allTweets, tweets...)
	}

	sort.Slice(allTweets, func(i, j int) bool {
		if allTweets[i].User == allTweets[j].User {
			return allTweets[i].PostedAt.After(allTweets[j].PostedAt)
		}
		return allTweets[i].User < allTweets[j].User
	})

	return allTweets, nil
}
