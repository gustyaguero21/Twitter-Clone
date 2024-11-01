package services

import (
	"twitter-clone/internal/models"
	"twitter-clone/internal/repository"
)

type FollowerServices struct {
	Repo *repository.Repository
}

func NewFollowerService(repo *repository.Repository) *FollowerServices {
	return &FollowerServices{Repo: repo}
}

func (fs *FollowerServices) FollowUser(follower, followed string) (models.Followers, error) {
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
