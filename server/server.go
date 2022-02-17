package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupServer() (router *gin.Engine) {
	router = gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))
	SetRoutes(router)
	return
}
