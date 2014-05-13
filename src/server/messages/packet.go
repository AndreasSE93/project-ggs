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
}

type HostNew struct {
	PacketID int `json:"PacketID"`
	RoomName string `json:"roomName"`
	MaxSize int `json:"maxSize"`
	GameName string `json:"GameName"`
}

type HostRoomPacket struct {
	PacketID int `json:"PacketID"`
	HostRoom ClientSection `json:"hostRoom"`
}

type RoomData struct {
	CS ClientSection
	SS ServerSection
}

type ClientSection struct {
	RoomID, MaxSize, ClientCount int
	RoomName, GameName string
	Clients []connection.Connector
}

type ServerSection struct {
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
	Rooms []ClientSection `json:"UserList"`
	Games []string `json:"GameHost"`
}

type ChatMessage struct {
	PacketID int `json:"PacketID"`
	Message string `json:"message"`
	User string `json:"user"`
}
