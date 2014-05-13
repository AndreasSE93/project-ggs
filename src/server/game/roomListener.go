package game

import(
	"fmt"
	"server/database/lobbyMap"
	"server/messages"
	"server/encoders"
)

type GameRoom struct {
	roomData messages.RoomData
	lm *lobbyMap.LobbyMap
}

func sendImmediateMessage(message string, cs messages.ClientSection) {
	fmt.Printf("Sending to %d clients: %s\n", cs.ClientCount, message)
	for i := 0; i < cs.ClientCount; i++ {
		cs.Clients[i].Connection.Write([]byte(message + "\n"))
	}
}

func gameRoomListener(gameRoom *GameRoom) {
	for {
		processed := <- gameRoom.roomData.SS.GameChan
		
		if processed.ID == messages.CHAT_ID {
			go sendImmediateMessage(encoders.EncodeChatMessage(messages.CHAT_ID, processed.ChatM, processed.Origin), gameRoom.roomData.CS)
		} else if processed.ID == messages.JOIN_ID {
			gameRoom.roomData = gameRoom.lm.GetRoom(gameRoom.roomData.CS.RoomID)
			go sendImmediateMessage(encoders.EncodeJoinedRoom(messages.JOIN_ID, &gameRoom.roomData), gameRoom.roomData.CS)
		}
	}
}

func CreateGameRoom(rd messages.RoomData, lm *lobbyMap.LobbyMap) {
	game := new(GameRoom)
	game.roomData = rd
	game.lm = lm

	go gameRoomListener(game)
}
