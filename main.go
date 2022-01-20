package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	SetEnv()
	SetRoutes(router)
	err := router.Run(":3000")
	if err != nil {
		return
	}
}