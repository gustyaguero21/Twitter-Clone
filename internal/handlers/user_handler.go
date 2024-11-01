package handlers

import (
	"net/http"
	"twitter-clone/internal/models"
	"twitter-clone/internal/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service services.Services
}

func NewUserHandler(service services.Services) *UserHandler {
	return &UserHandler{Service: service}
}

func (u *UserHandler) CreateUserHandler(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	user := models.Users{}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	createdUser, err := u.Service.CreateUser(ctx, user)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, createUserResponse(http.StatusOK, "User created successfully", createdUser))
}

func createUserResponse(status int, message string, user interface{}) models.UserResponse {
	return models.UserResponse{
		Status:  status,
		Message: message,
		User:    user,
	}
}
