package main

import "twitter-clone/internal/router"

func main() {
	router := router.StartRouter()

	router.Run("localhost:8080")
}
