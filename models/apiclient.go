package models

type APIAgent struct {
	Type    string `json:"type"`
	Status  string `json:"status"`
	Message ClientMessage
}

type ClientMessage struct {
	ID        uint32 `json:"id"`
	TimeStamp string `json:"timestamp"`
	Data      string `json:"data"`
}
