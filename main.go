package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mandrigin/gin-spa/spa"
)

func main() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))
	SetEnv()
	SetRoutes(router)
	router.Use(spa.Middleware("/", "./client/public")) // your build of React or other SPA
	err := router.Run(":8000")
	if err != nil {
		return
	}
}
