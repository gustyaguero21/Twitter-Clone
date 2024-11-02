package main

import (
	"fmt"
	"log"
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

	if err := router.Run(config.Port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
