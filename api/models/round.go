package models

var RoundModel *Round

type Round struct {
	PlayerId string `json:"id"`
	Question int    `json:"question"`
	Answer   int    `json:"answer"`
}