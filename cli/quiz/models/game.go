package models

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

func GetGameFromDTO(data interface{}) Game {
	m := data.(map[string]interface{})
	game := Game{}
	game.PlayerId = m["id"].(string)
	game.PlayerName = m["player"].(string)
	game.Questions = int(m["questions"].(float64))
	game.Answers = int(m["answers"].(float64))
	game.Pending = m["pending"].(bool)
	game.PendingId = int(m["pending_id"].(float64))
	game.Success = float32(m["success"].(float64))
	game.Position = int(m["position"].(float64))
	return game
}
