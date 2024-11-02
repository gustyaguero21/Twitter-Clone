package controllers

import (
	"net/http"
	"reflect"
	"twitter-clone/internal/models"
	"twitter-clone/internal/services"
	"twitter-clone/internal/utils"

	"github.com/gin-gonic/gin"
)

type FollowerHandler struct {
	Service services.FollowerServices
}

func NewFollowerHandler(service services.FollowerServices) *FollowerHandler {
	return &FollowerHandler{Service: service}
}

func (f *FollowerHandler) FollowUserHandler(ctx *gin.Context) {

	ctx.Header("Content-Type", "application/json")

	follower := ctx.Param("username")

	if follower == "" || reflect.TypeOf(follower).Kind() != reflect.String {
		ctx.JSON(400, gin.H{
			"error": "invalid query param",
		})
		return
	}

	followers := models.Followers{}

	if err := ctx.ShouldBindJSON(&followers); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err := f.Service.FollowUser(follower, followers.FollowingUsername)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	message := "Now you're following " + followers.FollowingUsername
	ctx.JSON(200, utils.CreateResponse(http.StatusOK, message))
}

func (f *FollowerHandler) Following(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	username := ctx.Param("username")

	if username == "" || reflect.TypeOf(username).Kind() != reflect.String {
		ctx.JSON(400, gin.H{
			"error": "invalid query param",
		})
		return
	}

	following_users, err := f.Service.ShowFollowers(username)

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
