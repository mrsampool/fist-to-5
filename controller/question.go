package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mrsampool/fist-to-5/db/model"
	"strconv"
)

func GetCurrentQuestion(c *gin.Context) {
	// TODO
}
func NewQuestion(c *gin.Context) {
	// TODO
}
func GetQuestionById(c *gin.Context) {
	questionIdParam := c.Param("questionId")
	questionIdInt, err := strconv.Atoi(questionIdParam)
	if err != nil {
		fmt.Println("Received invalid question ID param")
	}
	model.GetQuestionById(questionIdInt)
}
