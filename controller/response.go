package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mrsampool/fist-to-5/db/model/response"
	"net/http"
	"strconv"
)

func AddResponse(c *gin.Context) {
	questionIdParam := c.Param("questionId")
	questionIdInt, err := strconv.Atoi(questionIdParam)
	if err != nil {
		fmt.Println("Error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid question id"})
		return
	}
	var newResponse struct {
		StudentId int `json:"studentId"`
		Value     int `json:"value"`
		// QuestionId int `json:"questionId"`
	}
	err = c.BindJSON(&newResponse)
	if err != nil {
		fmt.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse request data"})
		return
	}
	result, err := response.InsertResponse(newResponse.StudentId, questionIdInt, newResponse.Value)
	if err != nil {
		fmt.Println("error: ", err.Error())
	}
	c.JSON(200, result)
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
