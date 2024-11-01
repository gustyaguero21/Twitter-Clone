package services

import (
	"context"
	"twitter-clone/internal/models"
	"twitter-clone/internal/repository"

	"github.com/google/uuid"
)

type Services struct {
	Repo *repository.Repository
}

func NewService(repo *repository.Repository) *Services {
	return &Services{Repo: repo}
}

func (s *Services) CreateUser(ctx context.Context, user models.Users) (models.Users, error) {
	newUser := models.Users{
		ID:       uuid.New(),
		Username: user.Username,
	}

	userRepo := repository.NewUserRepository(*s.Repo)

	if err := userRepo.SaveUser(newUser); err != nil {
		return models.Users{}, err
	}

	return newUser, nil
}
