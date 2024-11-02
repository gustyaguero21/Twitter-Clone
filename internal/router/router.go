package router

import "github.com/gin-gonic/gin"

func StartRouter() *gin.Engine {

	router := gin.Default()

	UrlMapping(router)

	return router
}
