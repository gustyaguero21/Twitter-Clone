package controllers

import (
	"net/http"
	"twitter-clone/internal/models"
	"twitter-clone/internal/services"
	"twitter-clone/internal/utils"

	"github.com/gin-gonic/gin"
)

type FollowerController struct {
	Service services.FollowerServices
}

func NewFollowerController(service services.FollowerServices) *FollowerController {
	return &FollowerController{Service: service}
}

func (f *FollowerController) FollowUserController(ctx *gin.Context) {

	ctx.Header("Content-Type", "application/json")

	follower := ctx.Param("username")

	followers := models.Followers{}

	if err := ctx.ShouldBindJSON(&followers); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err := f.Service.FollowUser(ctx, follower, followers.FollowingUsername)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	message := "Now you're following " + followers.FollowingUsername

	ctx.JSON(200, utils.CreateResponse(http.StatusOK, message))
}

func (f *FollowerController) Following(ctx *gin.Context) {

	ctx.Header("Content-Type", "application/json")

	username := ctx.Param("username")

	following_users, err := f.Service.ShowFollowers(ctx, username)

	if len(following_users) == 0 {
		ctx.JSON(200, gin.H{
			"message": "user aleady not following users",
		})
		return
	}

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, following_users)
}
