package messages

import (
	"server/database/lobbyMap"
)

const (
	CHAT_ID    = 100
	HOST_ID    = 101
	JOIN_ID    = 102
	REFRESH_ID = 103
)

type HostNew struct {
	PacketID int `json:"PacketID"`
	RoomName string `json:"roomName"`
	MaxSize int `json:"maxSize"`
	GameName string `json:"GameName"`
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
