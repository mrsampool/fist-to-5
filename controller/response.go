package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mrsampool/fist-to-5/db/model/response"
)

func AddResponse(c *gin.Context) {
	// TODO
}
func GetCurrentResponses(c *gin.Context) {
	responses, err := response.QueryResponsesByCurrentQuestion()
	if err != nil {
		fmt.Println("error: ", err.Error())
	}
	c.JSON(200, responses)
}
func GetResponsesByQuestion(c *gin.Context) {
	// TODO
}
