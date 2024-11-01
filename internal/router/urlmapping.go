package router

import (
	"log"
	"twitter-clone/internal/handlers"
	"twitter-clone/internal/repository"
	"twitter-clone/internal/services"

	"github.com/gin-gonic/gin"
)

func UrlMapping(r *gin.Engine) {

	repository, err := repository.NewRepository()
	if err != nil {
		log.Fatal("error initializing database.")
	}

	userService := services.NewService(&repository)
	userHandler := handlers.NewUserHandler(*userService)

	followerService := services.NewFollowerService(&repository)
	followerHandler := handlers.NewFollowerHandler(*followerService)

	router := r.Group("twitter-clone/api/v1")

	router.POST("/create-user", userHandler.CreateUserHandler)

	router.POST("/follow-user/:username", followerHandler.FollowUser)

}
