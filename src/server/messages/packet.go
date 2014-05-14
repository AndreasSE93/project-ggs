package messages

import (
	"time"
	"server/connection"
)
//-------------------------------------------------
//            ID FOR PACKAGES
const (
	PING_ID    = 0
	INIT_ID    = 99
	CHAT_ID    = 100
	HOST_ID    = 101
	JOIN_ID    = 102
	REFRESH_ID = 103
	TTT_CHAT_ID= 200
	TTT_MOVE_ID= 201
)

//-------------------------------------------------
//            PING

type Ping struct {
	PacketID int `json:"PacketID"`
	TimeStamp time.Time `json:"TimeStamp"`
	Payload interface{} `json:"Payload"`
}

//--------------------------------------------------
//           INCOMING FROM CLIENT

type ProcessedMessage struct {
	ID int `json:"PacketID"`
	Origin connection.Connector
	InitM InitMessage
	ChatM ChatMessage
	Host HostNew
	Join JoinExisting
	Update UpdateRooms
	MoveM MoveMessage
}

type InitMessage struct {
	PacketID int
	UserName string
}

type ChatMessage struct {
	PacketID int `json:"PacketID"`
	Message string `json:"message"`
	User string `json:"user"`
}

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

//------------------------------------------------
//             OUTGOING TO CLIENT

type HostRoomPacket struct {
	PacketID int `json:"PacketID"`
	HostRoom ClientSection `json:"hostRoom"`
	Player int `json:"Player"`
}


type RoomList struct {
	PacketID int `json:"PacketID"`
	Rooms []ClientSection `json:"UserList"`
	Games []string `json:"GameHost"`
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

//----------------------------------------------------
//              INSIDE OF LOBBYMAP

type RoomData struct {
	CS ClientSection
	SS ServerSection
}

type ClientSection struct {
	RoomID, MaxSize, ClientCount int
	RoomName, GameName, GameType string
	Clients []connection.Connector
}

type ServerSection struct {
	GameChan chan ProcessedMessage
}

//----------------------------------------------------

