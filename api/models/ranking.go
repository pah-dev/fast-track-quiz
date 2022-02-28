package models

var RankingModel *Ranking

type Ranking struct {
	PlayerId string  `json:"id"`
	Success  float32 `json:"success"`
}