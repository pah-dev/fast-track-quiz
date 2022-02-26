package models

type ApiResponse struct {
	Response string      `json:"response"`
	Data     interface{} `json:"data"`
}

type Question struct {
	Id       int    `json:"id"`
	Question string `json:"question"`
	Answer   int    `json:"answer"`
	Opt1     string `json:"1"`
	Opt2     string `json:"2"`
	Opt3     string `json:"3"`
	Opt4     string `json:"4"`
}

type Round struct {
	PlayerId string `json:"id"`
	Question int    `json:"question"`
	Answer   int    `json:"answer"`
}

type Game struct {
	PlayerId   string      `json:"id"`
	PlayerName string      `json:"player"`
	Questions  int         `json:"questions"`
	Answers    int         `json:"answers"`
	Pending    bool        `json:"pending"`
	PendingId  int         `json:"pending_id"`
	List       map[int]int `json:"list"`
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
	return game
}

func GetQuestionFromDTO(data interface{}) Question {
	m := data.(map[string]interface{})
	question := Question{}
	question.Id = int(m["id"].(float64))
	question.Question = m["question"].(string)
	question.Answer = int(m["answer"].(float64))
	question.Opt1 = m["1"].(string)
	question.Opt2 = m["2"].(string)
	question.Opt3 = m["3"].(string)
	question.Opt4 = m["4"].(string)
	return question
}

func GetRoundFromDTO(data interface{}) Round {
	m := data.(map[string]interface{})
	round := Round{}
	round.PlayerId = m["id"].(string)
	round.Question = int(m["question"].(float64))
	round.Answer = int(m["answer"].(float64))
	return round
}
