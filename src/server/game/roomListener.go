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
	roomData messages.RoomData
	lm *lobbyMap.LobbyMap
}

func sendImmediateMessage(message string, clients []connection.Connector) {
	for c := range clients {
		clients[c].Connection.Write([]byte(message + "\n"))
	}
}

func gameRoomListener(gameRoom *GameRoom) {
	for {
		processed := <- gameRoom.roomData.SS.GameChan
		
		if processed.ID == messages.CHAT_ID {
			go sendImmediateMessage(encoders.EncodeChatMessage(messages.CHAT_ID, processed.ChatM), gameRoom.roomData.CS.Clients)
		} else if processed.ID == messages.JOIN_ID {
			gameRoom.roomData = gameRoom.lm.GetRoom(gameRoom.roomData.CS.RoomID)
			go sendImmediateMessage(encoders.EncodeJoinedRoom(messages.JOIN_ID, &gameRoom.roomData), gameRoom.roomData.CS.Clients)
		}
	}
}

func CreateGameRoom(rd messages.RoomData, lm *lobbyMap.LobbyMap) {
	game := new(GameRoom)
	game.roomData = rd
	game.lm = lm

	go gameRoomListener(game)
}
