package models

var QuizModel *Quiz

type Quiz struct {
	Id       int    `json:"id"`
	Question string `json:"question"`
	Answer   int    `json:"answer"`
	Opt1     string `json:"1"`
	Opt2     string `json:"2"`
	Opt3     string `json:"3"`
	Opt4     string `json:"4"`
}
