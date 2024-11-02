package router

import (
	"log"
	"twitter-clone/internal/controllers"
	"twitter-clone/internal/repository"
	"twitter-clone/internal/services"

	"github.com/gin-gonic/gin"
)

func UrlMapping(r *gin.Engine) {

	repository, err := repository.NewRepository()
	if err != nil {
		log.Fatal("error initializing database.")
	}

	userService := services.NewUserService(&repository)
	userHandler := controllers.NewUserHandler(*userService)

	followerService := services.NewFollowerService(&repository)
	followerHandler := controllers.NewFollowerHandler(*followerService)

	postService := services.NewTweetService(&repository)
	postHandler := controllers.NewTweetHandler(*postService)

	router := r.Group("twitter-clone/api/v1")

	router.POST("/create-user", userHandler.CreateUserHandler)

	router.POST("/follow-user/:username", followerHandler.FollowUserHandler)

	router.GET("/followers/:username", followerHandler.Following)

	router.POST("/create-post/:username", postHandler.CreatePostHandler)

	router.GET("/timeline/:username", postHandler.TimelineHandler)

}
