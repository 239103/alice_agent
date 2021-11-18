package models

type APIClient struct {
	ID      uint64
	Type    string
	MsgBody ClientMessage
}

type ClientMessage struct {
}
