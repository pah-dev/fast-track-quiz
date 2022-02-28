package models

type Question struct {
	Id       int    `json:"id"`
	Question string `json:"question"`
	Answer   int    `json:"answer"`
	Opt1     string `json:"1"`
	Opt2     string `json:"2"`
	Opt3     string `json:"3"`
	Opt4     string `json:"4"`
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
