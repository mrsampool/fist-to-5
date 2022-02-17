package server

import (
	"github.com/gin-gonic/gin"
	"github.com/mandrigin/gin-spa/spa"
	"github.com/mrsampool/fist-to-5/controller"
)

func SetRoutes(router *gin.Engine) {
	apiRouter := router.Group("/api")
	{
		responsesRouter := apiRouter.Group("/responses")
		{
			responsesRouter.POST("/:questionId", controller.AddResponse)
			responsesRouter.GET("/:questionId", controller.GetResponsesByQuestion)
			responsesRouter.GET("/current", controller.GetCurrentResponses)
		}
		questionRouter := apiRouter.Group("/question")
		{
			questionRouter.GET("/:questionId", controller.GetQuestionById)
			questionRouter.GET("/current", controller.GetCurrentQuestion)
			questionRouter.POST("/", controller.PostQuestion)
		}
	}
	router.Use(spa.Middleware("/", "./client/public"))
}
