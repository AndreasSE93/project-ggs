package game

import(
//	"fmt"
//	"encoding/json"
	"server/database/lobbyMap"
	"server/messages"
	"server/encoders"
	"server/connection"
	"server/games"
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
		} else if processed.ID == messages.TTT_CHAT_ID {
			go sendImmediateMessage(encoders.EncodeChatMessage(messages.CHAT_ID, processed.ChatM), gameRoom.hostRoom.Clients)
		} else if processed.ID == messages.TTT_MOVE_ID {
			
			go sendImmediateMessage(encoders.EncodeMoveMessage(messages.TTT_MOVE_ID, processed.MoveM), gameRoom.hostRoom.Clients)
		
		}
	}
}



func ticTacToeListener (gameRoom *GameRoom) {
	gameBoard := games.InitBoard()
	
	for {
		processed := <- gameRoom.hostRoom.GameChan
		
		if processed.ID == messages.TTT_CHAT_ID {
			go sendImmediateMessage(encoders.EncodeChatMessage(messages.CHAT_ID, processed.ChatM), gameRoom.hostRoom.Clients)
		}  else if processed.ID == messages.JOIN_ID {
			gameRoom.hostRoom = gameRoom.lm.GetRoom(gameRoom.hostRoom.RoomID)
			go sendImmediateMessage(encoders.EncodeJoinedRoom(messages.JOIN_ID, &gameRoom.hostRoom), gameRoom.hostRoom.Clients)
		} else if processed.ID == messages.TTT_MOVE_ID {
			move := processed.MoveM.Move
			if games.IsValidMove(move, gameBoard) == 1 {
				gameBoard = games.MakeMove(move, processed.MoveM.Player, gameBoard)
				processed.MoveM.GameBoard = gameBoard
				processed.MoveM.HasWon = games.HasWon(gameBoard)
				processed.MoveM.IsDraw = games.IsDraw(gameBoard)
				processed.MoveM.IsValid = 1	
			} else {
				processed.MoveM.IsValid = 0
			}
			
			go sendImmediateMessage(encoders.EncodeMoveMessage(messages.TTT_MOVE_ID, processed.MoveM),  gameRoom.hostRoom.Clients)
			
		}
	}
}


func CreateGameRoom(hr messages.HostRoom, lm *lobbyMap.LobbyMap) {
	game := new(GameRoom)
	game.hostRoom = hr
	game.lm = lm
	
	switch hr.GameName {
	case "Dummy": go ticTacToeListener(game)
	
	default: go gameRoomListener(game)
	}
	
	

}
