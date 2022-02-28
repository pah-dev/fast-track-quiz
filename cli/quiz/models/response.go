package models

type ApiResponse struct {
	Response string      `json:"response"`
	Data     interface{} `json:"data"`
}
