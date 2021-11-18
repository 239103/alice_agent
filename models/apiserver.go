package models

type APIServer struct {
	ID      uint64
	Type    string
	MsgBody ServerMessage
}

type ServerMessage struct {
}
