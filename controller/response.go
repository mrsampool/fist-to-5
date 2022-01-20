package controller

import "github.com/gin-gonic/gin"

func AddResponse(c *gin.Context) {
	// TODO
}
func GetCurrentResponses(c *gin.Context) {
	c.JSON(200, gin.H{"handler": "GetCurrentResponses"})
}
func GetResponsesByQuestion(c *gin.Context) {
	// TODO
}
