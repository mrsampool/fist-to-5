package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mrsampool/fist-to-5/controller"
)

func SetRoutes(router *gin.Engine) {
	apiRouter := router.Group("/api")
	{
		responsesRouter := apiRouter.Group("/responses")
		{
			responsesRouter.POST("/:questionId", controller.AddResponse)
			responsesRouter.GET("/:questionId", controller.GetResponsesByQuestion)
			responsesRouter.GET("/", controller.GetCurrentResponses)
		}
		questionRouter := apiRouter.Group("/question")
		{
			questionRouter.GET("/current", controller.GetCurrentQuestion)
			questionRouter.GET("/:questionId", controller.GetQuestionById)
			questionRouter.POST("/", controller.PostQuestion)
		}
	}
}
