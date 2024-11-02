package controllers

import (
	"twitter-clone/internal/models"
	"twitter-clone/internal/services"

	"github.com/gin-gonic/gin"
)

type TweetController struct {
	Service services.TweetServices
}

func NewTweetController(service services.TweetServices) *TweetController {
	return &TweetController{Service: service}
}

func (t *TweetController) CreatePostController(ctx *gin.Context) {

	ctx.Header("Content-Type", "application/json")

	username := ctx.Param("username")

	tweet := models.Tweets{}

	if err := ctx.ShouldBindJSON(&tweet); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	createdPost, createdErr := t.Service.CreatePost(ctx, username, tweet.Content)

	if createdErr != nil {
		ctx.JSON(400, gin.H{
			"error": createdErr.Error(),
		})
		return
	}

	ctx.JSON(200, createdPost)
}

func (t *TweetController) TimelineController(ctx *gin.Context) {

	ctx.Header("Content-Type", "application/json")

	username := ctx.Param("username")

	getTimeline, getErr := t.Service.ShowTimeline(username)

	if getErr != nil {
		ctx.JSON(400, gin.H{
			"error": getErr.Error(),
		})
		return
	}
	if len(getTimeline) == 0 {
		ctx.JSON(200, gin.H{
			"message": "user currently not following any user",
		})
		return
	}

	if len(getTimeline) == 0 {
		ctx.JSON(400, gin.H{
			"message": "user currently not following any user",
		})
		return
	}

	ctx.JSON(200, getTimeline)
}
