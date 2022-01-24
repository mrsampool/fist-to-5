package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mrsampool/fist-to-5/db/model/response"
	"net/http"
	"strconv"
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
	questionIdParam := c.Param("questionId")
	questionIdInt, err := strconv.Atoi(questionIdParam)
	if err != nil {
		fmt.Println("Error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid question id"})
		return
	}
	responses, err := response.QueryResponsesByQuestionId(questionIdInt)
	if err != nil {
		fmt.Println("error: ", err.Error())
	}
	c.JSON(200, responses)
}
