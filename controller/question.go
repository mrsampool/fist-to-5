package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mrsampool/fist-to-5/db/model/question"
	"net/http"
)

type reqQuestion struct {
	QuestionText string `json:"text"`
}

func GetCurrentQuestion(c *gin.Context) {
	question, err := question.QueryCurrentQuestion()
	if err != nil {
		fmt.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, question)
}

func GetQuestionById(c *gin.Context) {
	/*
		questionIdParam := c.Param("questionId")
		questionIdInt, err := strconv.Atoi(questionIdParam)
		if err != nil {
			fmt.Println("Error: ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid question id"})
			return
		}
		question, err := model.QueryQuestionById(questionIdInt)
		if err != nil {
			fmt.Println("Error: ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "no question found with provided question id"})
			return
		}
		c.JSON(http.StatusOK, question)

	*/
}

func PostQuestion(c *gin.Context) {
	/*
		var question reqQuestion
		err := c.BindJSON(&question)
		if err != nil {
			fmt.Println("Error: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse request data"})
			return
		}
		if question.QuestionText == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid question data"})
			return
		}
		fmt.Println(question)
		newQuestion, err := model.InsertQuestion(question.QuestionText)
		if err != nil {
			fmt.Println("Error: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, newQuestion)

	*/
}
