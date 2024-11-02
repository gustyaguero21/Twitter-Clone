package controllers

import (
	"net/http"
	"twitter-clone/internal/models"
	"twitter-clone/internal/services"
	"twitter-clone/internal/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Service services.UserServices
}

func NewUserController(service services.UserServices) *UserController {
	return &UserController{Service: service}
}

func (u *UserController) CreateUserController(ctx *gin.Context) {

	ctx.Header("Content-Type", "application/json")

	user := models.Users{}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err := u.Service.CreateUser(ctx, user)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, utils.CreateResponse(http.StatusOK, "User created successfully"))
}
