package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mrsampool/fist-to-5/db/model/question"
	"net/http"
	"strconv"
)

func GetCurrentQuestion(c *gin.Context) {
	currQuestion, err := question.QueryCurrentQuestion()
	if err != nil {
		fmt.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, currQuestion)
}

func GetQuestionById(c *gin.Context) {
	questionIdParam := c.Param("questionId")
	questionIdInt, err := strconv.Atoi(questionIdParam)
	if err != nil {
		fmt.Println("Error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid question id"})
		return
	}
	foundQuestion, err := question.QueryById(questionIdInt)
	if err != nil {
		fmt.Println("Error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "no question found with provided question id"})
		return
	}
	c.JSON(http.StatusOK, foundQuestion)
}

func PostQuestion(c *gin.Context) {
	var newQuestion struct {
		Text string `json:"text"`
	}
	err := c.BindJSON(&newQuestion)
	if err != nil {
		fmt.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse request data"})
		return
	}
	if newQuestion.Text == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid question data"})
		return
	}
	result, err := question.Insert(newQuestion.Text)
	if err != nil {
		fmt.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
