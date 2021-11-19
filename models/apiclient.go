package models

import "time"

type APIClient struct {
	Type    string `json:"type"`
	Status	string `json:"status"`
	Message ClientMessage
}

type ClientMessage struct {
	ID        uint32    `json:"id"`
	TimeStamp time.Time `json:"timestamp"`
	Data      string    `json:"data"`
}
