package services

import (
	"context"
	"twitter-clone/internal/models"
	"twitter-clone/internal/repository"
)

type FollowerServices struct {
	Repo *repository.Repository
}

func NewFollowerService(repo *repository.Repository) *FollowerServices {
	return &FollowerServices{Repo: repo}
}

func (fs *FollowerServices) FollowUser(ctx context.Context, follower, followed string) (models.Followers, error) {

	newFollower := models.Followers{
		FollowerUsername:  follower,
		FollowingUsername: followed,
	}

	followRepo := repository.NewFollowerRepository(*fs.Repo)

	if err := followRepo.Follow(follower, followed); err != nil {
		return models.Followers{}, err
	}
	return newFollower, nil
}

func (fs *FollowerServices) ShowFollowers(ctx context.Context, username string) ([]string, error) {

	followRepo := repository.FollowerRepository(*fs.Repo)

	follower_list, err := followRepo.SearchFollowers(username)

	if err != nil {
		return nil, err
	}
	username_list := []string{}

	for _, follower := range follower_list {
		username_list = append(username_list, follower.FollowingUsername)
	}

	return username_list, nil
}
