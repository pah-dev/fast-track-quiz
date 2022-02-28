package models

type Round struct {
	PlayerId string `json:"id"`
	Question int    `json:"question"`
	Answer   int    `json:"answer"`
}

func GetRoundFromDTO(data interface{}) Round {
	m := data.(map[string]interface{})
	round := Round{}
	round.PlayerId = m["id"].(string)
	round.Question = int(m["question"].(float64))
	round.Answer = int(m["answer"].(float64))
	return round
}
