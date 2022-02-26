package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pah-dev/fast-track-quiz/api/controllers"
)

func Quiz(e *gin.Engine) {
	r := e.Group("/quiz")

	r.POST("/start", controllers.Quiz.StartQuiz)
	r.GET("/:id/question", controllers.Quiz.GetOneQuestion)
	r.POST("/answer", controllers.Quiz.AnswerQuestion)
	r.GET("/:id/end", controllers.Quiz.EndQuiz)

	r.GET("/questions", controllers.Quiz.GetQuestions)
	r.GET("/games", controllers.Quiz.GetGames)
}