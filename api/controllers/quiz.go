package controllers

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pah-dev/fast-track-quiz/api/database"
	"github.com/pah-dev/fast-track-quiz/api/models"
)

var Quiz *QuizController

type QuizController struct {
	
}

func (w *QuizController) StartQuiz(c *gin.Context){
	var game models.Game
	err := c.ShouldBindJSON(&game)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	game.PlayerId = uuid.New().String()
	game.List = map[int]int{}
	database.Games[game.PlayerId] = game
	c.IndentedJSON(http.StatusOK, gin.H{
		"response" : "OK",
		"data" : game,
	})
}

//GetQuizByID ... Get the quiz by id
func (w *QuizController) EndQuiz(c *gin.Context) {
	idPlayer := c.Params.ByName("id")
	game := database.Games[idPlayer]
	c.IndentedJSON(http.StatusOK, gin.H{
		"response" : "OK",
		"data" : game,
	})
}

//GetQuizByID ... Get the quiz by id
func (w *QuizController) GetOneQuestion(c *gin.Context) {
	playerId := c.Params.ByName("id")
	pend := PendingQuestion(playerId)
	if pend {
		c.IndentedJSON(http.StatusOK, gin.H{
			"response" : "error",
			"data" : "You have a pendging question",
		})
	}else{
		questionId := GetQuestionId(playerId)
		if questionId > 0 {
			c.IndentedJSON(http.StatusOK, gin.H{
				"response" : "OK",
				"data" : database.Questions[questionId-1],
			})
		}else{
			c.IndentedJSON(http.StatusOK, gin.H{
				"response" : "OK",
				"data" : "Quiz finished",
			})
		}
	}
}

//GetQuizByID ... Get the quiz by id
func (w *QuizController) AnswerQuestion(c *gin.Context) {
	var round models.Round
	err := c.ShouldBindJSON(&round)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	ans := CheckAnswer(round)
	if ans < 0 {
		c.IndentedJSON(http.StatusOK, gin.H{
			"response" : "error",
			"data" : "Wrong question ID",
		})
	}else if ans == 1 {
		c.IndentedJSON(http.StatusOK, gin.H{
			"response" : "OK",
			"data" : "Correct answer",
		})
	}else{
		c.IndentedJSON(http.StatusOK, gin.H{
			"response" : "OK",
			"data" : "Wrong answer",
		})
	}
}

//GetQuizByID ... Get the quiz by id
func (w *QuizController) GetQuestions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, database.Questions)
}

//GetQuizByID ... Get the quiz by id
func (w *QuizController) GetGames(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, database.Games)
}

func PendingQuestion(playerId string) bool{
	game := database.Games[playerId]
	return game.Pending
}

func GetQuestionId(playerId string) int{
	game := database.Games[playerId]
	if game.PlayerId != "" && len(game.List) < len(database.Questions) {
		if len(game.List) == 0 {
			questionId := rand.Intn(len(database.Questions))+1
			game.Pending = true
			game.PendingId = questionId
			game.Questions++
			game.List[questionId] = questionId
			database.Games[playerId] = game
			return questionId
		}
		i := -1
		for i < 0 && len(game.List) < len(database.Questions){
			questionId := rand.Intn(len(database.Questions))+1
			if game.List[questionId] == 0 {
				game.Pending = true
				game.PendingId = questionId
				game.Questions++
				game.List[questionId] = questionId
				database.Games[playerId] = game
				return questionId
			}
		}
	}
	return -1
}

func CheckAnswer(round models.Round) int{
	resp := -1
	quiz := database.Questions[round.Question-1]
	player := database.Games[round.PlayerId]
	if player.PendingId == round.Question {
		player.Pending = false
		player.PendingId = 0
		if quiz.Answer == round.Answer {
			player.Answers++
			resp = 1
		}else{
			resp = 0
		}
		database.Games[round.PlayerId] = player
	}
	return resp
}

// //GetQuizByID ... Get the quiz by id
// func (w *QuizController) GetQuizByID(c *gin.Context) {
// 	var quiz models.Quiz
// 	id := c.Params.ByName("id")
// 	db := c.MustGet("db").(*gorm.DB)
// 	err := models.QuizModel.GetQuizByID(db, &quiz, id)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.JSON(http.StatusOK, quiz)
// 	}
// }

// func (w *QuizController) Balance(c *gin.Context){
// 	var quiz models.Quiz
// 	id := c.Params.ByName("id")
// 	db := c.MustGet("db").(*gorm.DB)
// 	err := models.QuizModel.GetQuizByID(db, &quiz, id)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.JSON(http.StatusOK, quiz)
// 	}
// }

// func (w *QuizController) Credit(c *gin.Context){
// 	var quiz models.UpdateQuiz
// 	var oldQuiz models.Quiz
// 	id := c.Params.ByName("id")
// 	db := c.MustGet("db").(*gorm.DB)
// 	err := c.ShouldBindJSON(&quiz)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, err)
// 		return
// 	}
// 	err = models.QuizModel.GetQuizByID(db, &oldQuiz, id)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	}
// 	err = models.QuizModel.CreditQuiz(db, &oldQuiz, &quiz)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	}
// 	c.JSON(http.StatusOK, oldQuiz)
// }

// func (w *QuizController) Debit(c *gin.Context){
// 	var quiz models.UpdateQuiz
// 	var oldQuiz models.Quiz
// 	id := c.Params.ByName("id")
// 	db := c.MustGet("db").(*gorm.DB)
// 	err := c.ShouldBindJSON(&quiz)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, err)
// 		return
// 	}
// 	err = models.QuizModel.GetQuizByID(db, &oldQuiz, id)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	}
// 	err = models.QuizModel.DebitQuiz(db, &oldQuiz, &quiz)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	}
// 	c.JSON(http.StatusOK, oldQuiz)
// }