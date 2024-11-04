package services

import (
	"context"
	"twitter-clone/internal/models"
	"twitter-clone/internal/repository"

	"github.com/google/uuid"
)

type UserServices struct {
	Repo *repository.Repository
}

func NewUserService(repo *repository.Repository) *UserServices {
	return &UserServices{Repo: repo}
}

func (us *UserServices) CreateUser(ctx context.Context, user models.Users) (models.Users, error) {

	newUser := models.Users{
		ID:       uuid.New(),
		Username: user.Username,
	}

	userRepo := repository.NewUserRepository(*us.Repo)

	if err := userRepo.SaveUser(newUser); err != nil {
		return models.Users{}, err
	}

	return newUser, nil
}
