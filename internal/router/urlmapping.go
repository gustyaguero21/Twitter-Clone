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
	userHandler := controllers.NewUserController(*userService)

	followerService := services.NewFollowerService(&repository)
	followerHandler := controllers.NewFollowerController(*followerService)

	postService := services.NewTweetService(&repository)
	postHandler := controllers.NewTweetController(*postService)

	router := r.Group("twitter-clone/api/v1")

	router.POST("/create-user", userHandler.CreateUserController)

	router.POST("/follow-user/:username", followerHandler.FollowUserController)

	router.GET("/followers/:username", followerHandler.Following)

	router.POST("/create-tweet/:username", postHandler.CreatePostController)

	router.GET("/timeline/:username", postHandler.TimelineController)

}
