package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))
	SetEnv()
	SetRoutes(router)
	err := router.Run(":8000")
	if err != nil {
		return
	}
}
