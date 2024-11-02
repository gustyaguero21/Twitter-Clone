package controllers

import (
	"twitter-clone/internal/models"
	"twitter-clone/internal/services"

	"github.com/gin-gonic/gin"
)

type TweetHandler struct {
	Service services.TweetServices
}

func NewTweetHandler(service services.TweetServices) *TweetHandler {
	return &TweetHandler{Service: service}
}

func (t *TweetHandler) CreatePostHandler(ctx *gin.Context) {
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
