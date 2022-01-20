package controller

import "github.com/mrsampool/fist-to-5/db/model"
import "github.com/gin-gonic/gin"

func GetCurrentQuestion(c *gin.Context) {
	// TODO
}
func NewQuestion(c *gin.Context) {
	// TODO
}
func GetAllQuestions(c *gin.Context) {
	model.GetAllQuestions()
}
