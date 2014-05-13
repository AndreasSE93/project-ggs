package messages

import (
	"time"
	"server/connection"
)

const (
	PING_ID    = 0
	CHAT_ID    = 100
	HOST_ID    = 101
	JOIN_ID    = 102
	REFRESH_ID = 103
	TTT_CHAT_ID= 200
	TTT_MOVE_ID= 201
)

type Ping struct {
	PacketID int `json:"PacketID"`
	TimeStamp time.Time `json:"TimeStamp"`
	Payload interface{} `json:"Payload"`
}

type ProcessedMessage struct {
	ID int `json:"PacketID"`
	ChatM ChatMessage
	Host HostNew
	Join JoinExisting
	Update UpdateRooms
	MoveM MoveMessage
}

type HostNew struct {
	PacketID int `json:"PacketID"`
	RoomName string `json:"roomName"`
	MaxSize int `json:"maxSize"`
	GameName string `json:"GameName"`
}

type HostRoomPacket struct {
	PacketID int `json:"PacketID"`
	HostRoom HostRoom `json:"hostRoom"`
}

type RoomInfo struct {
	RoomID, MaxSize, ClientCount int
	RoomName, GameName string
}

type HostRoom struct {
	RoomID, MaxSize, ClientCount int
	RoomName, GameName string
	Clients []connection.Connector
	GameChan chan ProcessedMessage
}

type JoinExisting struct {
	PacketID int `json:"PacketID"`
	RoomID int `json:"RoomID"`
}

type UpdateRooms struct {
	PacketID int `json:"PacketID"`
}

type RoomList struct {
	PacketID int `json:"PacketID"`
	Rooms []RoomInfo `json:"UserList"`
	Games []string `json:"GameHost"`
}

type ChatMessage struct {
	PacketID int `json:"PacketID"`
	Message string `json:"message"`
	User string `json:"user"`
}

type MoveMessage struct {
	PacketID int `json:"PacketID"`
	GameBoard []int  `json:"GameBoard"`
	Move int `json:"Move"`
	IsDraw int `json:"IsDraw"`
	HasWon int `json:"HasWon"`
	Player int `json:"Player"`
	IsValid int `json:"IsValid"`
	
}
