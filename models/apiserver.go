package models

type APIServer struct {
	ID      uint32
	Type    string
	MsgBody ServerMessage
}

type ServerMessage struct {
}
