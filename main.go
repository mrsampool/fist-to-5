package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router, err := setupServer()
	err = router.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func setupServer() (router *gin.Engine, err error) {
	router = gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))
	err = SetEnv()
	if err != nil {
		return
	}
	SetRoutes(router)
	return
}
