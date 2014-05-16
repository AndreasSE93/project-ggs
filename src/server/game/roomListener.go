package game

import(
	"fmt"
	"server/database/lobbyMap"
	"server/messages"
	"server/encoders"
	"server/games"
	"time"
)

type GameRoom struct {
	roomData messages.RoomData
	lm *lobbyMap.LobbyMap
}

func sendImmediateMessage(message string, cs messages.ClientSection) {
	//fmt.Printf("Sending to %d clients: %s\n", cs.ClientCount, message)
	for i := 0; i < cs.ClientCount; i++ {
		cs.Clients[i].Connection.Write([]byte(message + "\n"))
		//fmt.Println(cs.ClientCount)
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

func ticTacToeListener (gameRoom *GameRoom) {
	gameBoard := games.InitBoard()
	
	for {
		processed := <- gameRoom.roomData.SS.GameChan
		
		if processed.ID == messages.CHAT_ID {
			go sendImmediateMessage(encoders.EncodeChatMessage(messages.CHAT_ID, processed.ChatM, processed.Origin), gameRoom.roomData.CS)
		}  else if processed.ID == messages.JOIN_ID {
			gameRoom.roomData = gameRoom.lm.GetRoom(gameRoom.roomData.CS.RoomID)
			go sendImmediateMessage(encoders.EncodeJoinedRoom(messages.JOIN_ID, &gameRoom.roomData), gameRoom.roomData.CS)
		} else if processed.ID == messages.TTT_MOVE_ID {
			move := processed.MoveM.Move
			if games.IsValidMove(move, gameBoard) == 1 {
				gameBoard = games.MakeMove(move, processed.MoveM.Player, gameBoard)
				processed.MoveM.GameBoard = gameBoard
				processed.MoveM.HasWon = games.HasWon(gameBoard)
				
				processed.MoveM.IsDraw = games.IsDraw(gameBoard)
				if (processed.MoveM.IsDraw == 1 || processed.MoveM.HasWon != 0){
					gameBoard = games.ClearBoard(gameBoard)
				}
				processed.MoveM.IsValid = 1	
			} else {
				processed.MoveM.IsValid = 0
			}
			
			go sendImmediateMessage(encoders.EncodeMoveMessage(messages.TTT_MOVE_ID, processed.MoveM),  gameRoom.roomData.CS)
			

		}
	}
}


func snakeListener(gameRoom *GameRoom){
	
	PlayerArray := []messages.Player{games.MakePlayerSnakes(1),games.MakePlayerSnakes(2),games.MakePlayerSnakes(3),games.MakePlayerSnakes(4)}
	go snakesHandler(&PlayerArray, gameRoom)
	for {
		processed := <- gameRoom.roomData.SS.GameChan
		fmt.Println(gameRoom.roomData.CS.Clients)
		if processed.ID == messages.SNAKES_CLIENT_ID {
			fmt.Println(processed.Snakes.Move, "hhhhhhhhhhhhhhhhhhhhhheeeeeeeeeeeeeeeee")
			PlayerArray[processed.Snakes.PlayerID-1] = games.UpdateMoveSnakes(PlayerArray[processed.Snakes.PlayerID-1], processed.Snakes.Move)
		}
	}

}


func snakesHandler(pA *[]messages.Player, gameRoom *GameRoom ){
	gameBoard := games.InitBoardSnakes()
	
	for{
		gameRoom.roomData = gameRoom.lm.GetRoom(gameRoom.roomData.CS.RoomID)
		*pA = games.UpdateAllMovesSnakes( *pA, gameBoard)
		gameBoard = games.DoMove(*pA, gameBoard)
		go sendImmediateMessage(encoders.EncodeSnakeMessage(messages.SNAKES_MOVES_ID, *pA), gameRoom.roomData.CS)
		fmt.Println(gameRoom.roomData.CS.ClientCount) 
		time.Sleep(5 * time.Millisecond )
	}
	

}

func CreateGameRoom(rd messages.RoomData, lm *lobbyMap.LobbyMap) {

	game := new(GameRoom)
	game.roomData = rd
	game.lm = lm
	
	switch rd.CS.GameName {
	case "Dummy": go ticTacToeListener(game)

	case "Snakes": go snakeListener(game)
	
	default: go gameRoomListener(game)
	}
	
	

}
