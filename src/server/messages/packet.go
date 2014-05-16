package messages

import (
	"time"
	"server/connection"
)
//-------------------------------------------------
//            ID FOR PACKAGES
const (
	PING_ID      = 0
	INIT_ID      = 99
	CHAT_ID      = 100
	HOST_ID      = 101
	JOIN_ID      = 102
	REFRESH_ID   = 103
	KICK_ID      = 404
	TTT_CHAT_ID  = 200
	TTT_MOVE_ID  = 201
	STARTABLE_ID = 202
	START_ID     = 203
	SNAKES_CLIENT_ID = 301
	SNAKES_MOVES_ID = 302
	
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
	Snakes SnakesEvent
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

type KickMessage struct {
	PacketID int `json:"PacketID"`
}

//------------------------------------------------
//             OUTGOING TO CLIENT

type HostRoomPacket struct {
	PacketID int `json:"PacketID"`
	HostRoom ClientSection `json:"hostRoom"`
}

type Leave struct {
	PacketID int `json:"PacketID"`
}

type RoomList struct {
	PacketID int `json:"PacketID"`
	Rooms []ClientSection `json:"UserList"`
	Games []string `json:"GameHost"`
}

type Startable struct {
	PacketID int `json:"PacketID"`
	IsStartable bool `json:"IsStartable"`
}

type Started struct {
	PacketID int `json:"PacketID"`
	Started bool `json:"Started"`
	Player int `json:"Player"`
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
//                ACHTUNG


type SnakesEvent struct {
	PacketID int `json:"PacketID"`
	PlayerID int `json:"PlayerID"`
	Move string `json:"KeyEvent"`

}



type Player struct{

	PlayerName string `json:"PlayerName"`
	PlayerID int `json:"PlayerID"`
	PosX float64 `json:"PosX"`
	PosY float64 `json:"PosY"`
	Alive bool `json:"Alive"`
	Direction float64 

}

type SnakeMessage struct {
	PacketID int `json:"PacketID"`
	PlayerArray []Player `json:"PlayerArray"`


}

