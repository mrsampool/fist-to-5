package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mrsampool/fist-to-5/controller"
)

func SetRoutes(router *gin.Engine) {
	router.GET("/api/responses", controller.GetCurrentResponses)
	router.POST("/api/responses/:questionId", controller.AddResponse)
	router.GET("/api/responses/:questionId", controller.GetResponsesByQuestion)
	router.GET("/api/question/current", controller.GetCurrentQuestion)
	router.GET("/api/question/:questionId", controller.GetQuestionById)
	router.POST("/api/question", controller.PostQuestion)
}
