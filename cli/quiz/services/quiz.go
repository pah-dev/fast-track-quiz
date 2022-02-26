package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pah-dev/fast-track-quiz/cli/quiz/models"
	"github.com/pah-dev/fast-track-quiz/cli/quiz/utils"
)

//	StartQuiz ... Insert new game
func StartQuiz(name string) (err error) {
	apiUrl := utils.ConfigKeyValuePairGet("api_url")
	game := models.Game{}
	game.PlayerName = name
	body, _ := json.Marshal(game)
	postBody := bytes.NewBuffer(body)
	resp, err := http.Post(fmt.Sprint(apiUrl)+"/quiz/start","application/json", postBody)
	if err != nil {
		utils.PrintError(fmt.Sprint(err))
	}
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		utils.PrintError(fmt.Sprint(err))
	}
	var response models.ApiResponse
    json.Unmarshal(responseBody, &response)
	if response.Response == "OK"{
		fmt.Println(response.Data)
		game := models.GetGameFromDTO(response.Data)
		utils.ConfigKeyValuePairUpdate("player_id", game.PlayerId)
		fmt.Println("")
		fmt.Println("###############################################")
		fmt.Println("		WELCOME " + fmt.Sprint(game.PlayerName))
		fmt.Println("  id: " + fmt.Sprint(game.PlayerId))
		fmt.Println("###############################################")
		fmt.Println("")
		fmt.Println("  Type quiz qn to get a question")
		fmt.Println("")
		fmt.Println("###############################################")
		fmt.Println("")
	}else{
		utils.PrintError(fmt.Sprint(response.Data))
	}
	
	return
}

//  GetOneQuestion to answer
func GetOneQuestion() (err error) {
	apiUrl := utils.ConfigKeyValuePairGet("api_url")
	playerId := utils.ConfigKeyValuePairGet("player_id")
	if playerId == nil {
		utils.PrintError("Type quiz init to start the game")
		return
	}
	resp, err := http.Get(fmt.Sprint(apiUrl)+"/quiz/"+fmt.Sprint(playerId)+"/question")
	if err != nil {
		utils.PrintError(fmt.Sprint(err))
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		utils.PrintError(fmt.Sprint(err))
	}
	var response models.ApiResponse
    json.Unmarshal(body, &response)
	if response.Response == "OK"{
		question := models.GetQuestionFromDTO(response.Data)
		fmt.Println(question)
		fmt.Println("")
		fmt.Println("###############################################")
		fmt.Println("		QUESTION ID: " + fmt.Sprint(question.Id))
		fmt.Println("###############################################")
		fmt.Println(question.Question)
		fmt.Println("1- " + fmt.Sprint(question.Opt1))
		fmt.Println("2- " + fmt.Sprint(question.Opt2))
		fmt.Println("3- " + fmt.Sprint(question.Opt3))
		fmt.Println("4- " + fmt.Sprint(question.Opt4))
		fmt.Println("-----------------------------------------------")
		fmt.Println("")
		fmt.Println("Type quiz an with Question ID and Answer number")	
		fmt.Println("Type quiz an -h for help")
		fmt.Println("###############################################")
		fmt.Println("")
	}else{		
		utils.PrintError(fmt.Sprint(response.Data))
	}
	return
}

func AnswerQuestion(question int, answer int)(err error) {
	apiUrl := utils.ConfigKeyValuePairGet("api_url")
	playerId := utils.ConfigKeyValuePairGet("player_id")
	if playerId == nil {
		utils.PrintError("Type quiz init to start the game")
		return
	}
	round := models.Round{}
	round.PlayerId = playerId.(string)
	round.Answer = answer
	round.Question = question
	body, _ := json.Marshal(round)
	postBody := bytes.NewBuffer(body)
	resp, err := http.Post(fmt.Sprint(apiUrl)+"/quiz/answer","application/json", postBody)
	if err != nil {
		utils.PrintError(fmt.Sprint(err))
	}
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		utils.PrintError(fmt.Sprint(err))
	}
	var response models.ApiResponse
    json.Unmarshal(responseBody, &response)
	if response.Response == "OK"{
		fmt.Println("")
		fmt.Println("###############################################")
		fmt.Println("??-------------------------------------------??")
		fmt.Println("		" + fmt.Sprint(response.Data))
		fmt.Println("??-------------------------------------------??")
		fmt.Println("")
		fmt.Println("Type quiz qn to get another question")
		fmt.Println("Type quiz en to finish and see your score")
		fmt.Println("###############################################")
		fmt.Println("")
	}else{
		utils.PrintError(fmt.Sprint(response.Data))
	}
	return
}

//  EndQuiz and get score
func EndQuiz()(err error){
	apiUrl := utils.ConfigKeyValuePairGet("api_url")
	playerId := utils.ConfigKeyValuePairGet("player_id")
	if playerId == nil {
		utils.PrintError("Type quiz init to start the game")
		return
	}
	resp, err := http.Get(fmt.Sprint(apiUrl)+"/quiz/"+fmt.Sprint(playerId)+"/end")
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var response models.ApiResponse
    json.Unmarshal(body, &response)
	if response.Response == "OK"{
		game := models.GetGameFromDTO(response.Data)
		fmt.Println("")
		fmt.Println("###############################################")
		fmt.Println("		QUIZ FINISHED")
		fmt.Println("###############################################")
		fmt.Println(fmt.Sprint(game.PlayerName) + ", this is your result:")
		fmt.Println("  Total answers:     " + fmt.Sprint(game.Questions))
		fmt.Println("  Correct answers:   " + fmt.Sprint(game.Answers))
		percent := (float32(game.Answers) / float32(game.Questions)) * 100
		fmt.Println("-----------------------------------------------")
		fmt.Println("  " + fmt.Sprint(percent) + "%" + " success")
		fmt.Println("###############################################")
		fmt.Println("")

		utils.DeleteKeyHack("player_id")
	}else{
		utils.PrintError(fmt.Sprint(response.Data))
	}
	return
}