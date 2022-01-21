package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mrsampool/fist-to-5/db/model"
	"net/http"
	"strconv"
)

func GetCurrentQuestion(c *gin.Context) {
	question, err := model.GetCurrentQuestion()
	if err != nil {
		fmt.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, question)
}

func NewQuestion(c *gin.Context) {
	// TODO
}

func GetQuestionById(c *gin.Context) {
	questionIdParam := c.Param("questionId")
	questionIdInt, err := strconv.Atoi(questionIdParam)
	if err != nil {
		fmt.Println("Error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid question id"})
		return
	}
	question, err := model.GetQuestionById(questionIdInt)
	if err != nil {
		fmt.Println("Error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "no question found with provided question id"})
		return
	}
	c.JSON(http.StatusOK, question)
}
