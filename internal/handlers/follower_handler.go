package handlers

import (
	"net/http"
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

func (f *FollowerHandler) FollowUser(ctx *gin.Context) {

	ctx.Header("Content-Type", "application/json")

	follower := ctx.Param("username")

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