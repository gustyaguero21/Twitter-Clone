package main

import (
	"fmt"
	"twitter-clone/cmd/config"
	"twitter-clone/internal/repository"
	"twitter-clone/internal/router"
)

func main() {

	_, err := repository.NewRepository()

	if err != nil {
		fmt.Println("Error initializing db")
	}

	router := router.StartRouter()

	router.Run(config.Port)
}
