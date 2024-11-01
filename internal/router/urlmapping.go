package router

import "github.com/gin-gonic/gin"

func UrlMapping(r *gin.Engine) {
	router := r.Group("twitter-clone/api/v1")

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
