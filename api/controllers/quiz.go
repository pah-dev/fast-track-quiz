package controllers

import (
	"math/rand"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pah-dev/fast-track-quiz/api/database"
	"github.com/pah-dev/fast-track-quiz/api/models"
	"github.com/pah-dev/fast-track-quiz/api/utils"
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

func (w *QuizController) EndQuiz(c *gin.Context) {
	idPlayer := c.Params.ByName("id")
	game := database.Games[idPlayer]
	percent := (float32(game.Answers) / float32(game.Questions)) * 100
	game.Success = utils.RoundUp(percent)
	rank := models.Ranking{PlayerId: idPlayer, Success: percent}
	database.Ranking = append(database.Ranking, rank)
	pos := getPosition(idPlayer)
	game.Position = pos
	database.Games[idPlayer] = game
	c.IndentedJSON(http.StatusOK, gin.H{
		"response" : "OK",
		"data" : game,
	})
}

func (w *QuizController) GetOneQuestion(c *gin.Context) {
	playerId := c.Params.ByName("id")
	pend := pendingQuestion(playerId)
	if pend {
		c.IndentedJSON(http.StatusOK, gin.H{
			"response" : "error",
			"data" : "You have a pending question",
		})
	}else{
		questionId := getQuestionId(playerId)
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

func (w *QuizController) AnswerQuestion(c *gin.Context) {
	var round models.Round
	err := c.ShouldBindJSON(&round)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	ans := checkAnswer(round)
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

func (w *QuizController) GetQuestions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, database.Questions)
}

func (w *QuizController) GetGames(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, database.Games)
}

func (w *QuizController) GetRanking(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, database.Ranking)
}

func pendingQuestion(playerId string) bool{
	game := database.Games[playerId]
	return game.Pending
}

// Function to get a random question and not repeated
func getQuestionId(playerId string) int{
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

func checkAnswer(round models.Round) int{
	resp := -1
	if(round.Question > len(database.Questions)){
		return resp
	}
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

func getPosition(playerId string) int{
	pos := len(database.Ranking)
	sort.SliceStable(database.Ranking, func(i, j int) bool {
		return database.Ranking[i].Success < database.Ranking[j].Success
	})
	for i := 0; i < len(database.Ranking); i++ {
		if database.Ranking[i].PlayerId == playerId {
			pos = i+1
		}
	}
	rank := (pos / len(database.Ranking))*100
	return rank
}
