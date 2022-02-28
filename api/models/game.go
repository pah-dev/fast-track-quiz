package models

var GameModel *Game

type Game struct {
	PlayerId   string      `json:"id"`
	PlayerName string      `json:"player"`
	Questions  int         `json:"questions"`
	Answers    int         `json:"answers"`
	Pending    bool        `json:"pending"`
	PendingId  int         `json:"pending_id"`
	List       map[int]int `json:"list"`
	Success    float32     `json:"success"`
	Position   int         `json:"position"`
}
