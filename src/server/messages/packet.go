package messages

import (
	"time"
	"server/database/lobbyMap"
)

const (
	PING_ID    = 0
	CHAT_ID    = 100
	HOST_ID    = 101
	JOIN_ID    = 102
	REFRESH_ID = 103
)

type Ping struct {
	PacketID int `json:"PacketID"`
	TimeStamp time.Time `json:"TimeStamp"`
	Payload interface{} `json:"Payload"`
}

type HostNew struct {
	PacketID int `json:"PacketID"`
	RoomName string `json:"roomName"`
	MaxSize int `json:"maxSize"`
	GameName string `json:"GameName"`
}

type HostRoom struct {
	PacketID int `json:"PacketID"`
	HostRoom lobbyMap.HostRoom `json:"PacketID"`
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
	Rooms []lobbyMap.HostRoom `json:"UserList"`
	Games []string `json:"GameHost"`
}

type ChatMessage struct {
	PacketID int `json:"PacketID"`
	Message string `json:"message"`
	User string `json:"user"`
}
