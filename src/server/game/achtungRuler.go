package game

import(
	"time"
	"server/encoders"
	"server/games"
	"server/messages"
)

func snakeListener(gameRoom *GameRoom) {
	for {
		processed := <- gameRoom.RuleChan
		if processed.ID == messages.START_ID {
			break
		}
	}
	newGameRoom := make(chan messages.RoomData)
	termChan    := make(chan interface{})
	PlayerArray := games.InitAchtungPlayerArray(4)
	go snakesHandler(PlayerArray, gameRoom, newGameRoom, termChan)
	for {
		processed := <- gameRoom.RuleChan
		if processed.ID == messages.SNAKES_CLIENT_ID {
			PlayerArray[processed.Snakes.PlayerID-1] = games.UpdateMoveSnakes(PlayerArray[processed.Snakes.PlayerID-1], processed.Snakes.Move)
		} else if processed.ID == messages.JOIN_ID {
			newGameRoom <- gameRoom.lm.GetRoom(gameRoom.roomData.CS.RoomID)
		} else if processed.ID == messages.ROOM_CLOSED_ID {
			termChan <- nil
		}

	}
}

func snakesHandler(pA []messages.Player, gameRoom *GameRoom, newGameRoom chan messages.RoomData, termChan chan interface{}) { 
	gameBoard := games.InitBoardSnakes()
	
	var Time uint16
	for {
		select {
		case newerGameRoom := <- newGameRoom:
			gameRoom.roomData = newerGameRoom
		case <- termChan:
			return
		default:
		}
		
		//gameRoom.roomData = gameRoom.lm.GetRoom(gameRoom.roomData.CS.RoomID)
		pA = games.UpdateAllMovesSnakes(pA, gameBoard, Time)
		games.DoMove(pA, gameBoard, Time)
		clear := games.AchtungFinished(pA, gameBoard)
		gameRoom.SendMult <- MultipleMessage{encoders.EncodeSnakeMessage(messages.SNAKES_MOVES_ID, pA, clear), gameRoom.roomData.CS.ClientCount, gameRoom.roomData.CS.Clients}
		if clear {
			Time = 0
			time.Sleep(3 * time.Second)
		} else {
			time.Sleep(30 * time.Millisecond )
		}
		Time++
	}
}
