package game

import(
//	"fmt"
//	"encoding/json"
	"server/database/lobbyMap"
	"server/messages"
	"server/encoders"
	"server/connection"
)

type GameRoom struct {
	hostRoom messages.HostRoom
	lm *lobbyMap.LobbyMap
}

func sendImmediateMessage(message string, clients []connection.Connector) {
	for c := range clients {
		clients[c].Connection.Write([]byte(message + "\n"))
	}
}

func gameRoomListener(gameRoom *GameRoom) {
	for {
		processed := <- gameRoom.hostRoom.GameChan
		
		if processed.ID == messages.CHAT_ID {
			go sendImmediateMessage(encoders.EncodeChatMessage(messages.CHAT_ID, processed.ChatM), gameRoom.hostRoom.Clients)
		} else if processed.ID == messages.JOIN_ID {
			gameRoom.hostRoom = gameRoom.lm.GetRoom(gameRoom.hostRoom.RoomID)
			go sendImmediateMessage(encoders.EncodeJoinedRoom(messages.JOIN_ID, &gameRoom.hostRoom), gameRoom.hostRoom.Clients)
		}
	}
}

func CreateGameRoom(hr messages.HostRoom, lm *lobbyMap.LobbyMap) {
	game := new(GameRoom)
	game.hostRoom = hr
	game.lm = lm

	go gameRoomListener(game)
}
