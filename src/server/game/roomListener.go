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
	GameType string
	RuleChan chan messages.ProcessedMessage
}

func sendImmediateMessage(message string, cs messages.ClientSection) {
	fmt.Printf("Sending to %d clients: %s\n", cs.ClientCount, message)
	for i := 0; i < cs.ClientCount; i++ {
		cs.Clients[i].Connection.Write([]byte(message + "\n"))
	}
}

func gameRoomListener(gameRoom *GameRoom) {
	go sendImmediateMessage(encoders.EncodeHostedRoom(messages.HOST_ID, gameRoom.roomData),gameRoom.roomData.CS)
	go sendImmediateMessage(encoders.EncodeHostedRoom(messages.HOST_ID, gameRoom.roomData),gameRoom.roomData.CS)
	for {
		processed := <- gameRoom.roomData.SS.GameChan
		
		if processed.ID == messages.CHAT_ID {
			go sendImmediateMessage(encoders.EncodeChatMessage(messages.CHAT_ID, processed.ChatM, processed.Origin), gameRoom.roomData.CS)
		} else if processed.ID == messages.JOIN_ID {
			gameRoom.roomData = gameRoom.lm.GetRoom(gameRoom.roomData.CS.RoomID)
			go sendImmediateMessage(encoders.EncodeJoinedRoom(messages.JOIN_ID, &gameRoom.roomData), gameRoom.roomData.CS)
			go sendImmediateMessage(encoders.EncodeJoinedRoom(messages.JOIN_ID, &gameRoom.roomData), gameRoom.roomData.CS)
		} else if processed.ID == messages.TTT_MOVE_ID {
			gameRoom.RuleChan <- processed
		}
	}
}

func CreateGameRoom(rd messages.RoomData, lm *lobbyMap.LobbyMap) {

	game := new(GameRoom)
	game.roomData = rd
	game.lm = lm
	game.RuleChan = make(chan messages.ProcessedMessage)
	game.GameType = rd.CS.GameName	
	
	switch game.GameType {
	case "TicTacToe": 
		go InitTicTac(game)
		go gameRoomListener(game)

	case "Achtung":
		go InitAchtung(game)
		go gameRoomListener(game)
		
	default: go gameRoomListener(game)
	}
}
