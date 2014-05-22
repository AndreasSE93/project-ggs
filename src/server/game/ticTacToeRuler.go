package game

import(
	"server/messages"
	"server/games"
	"server/encoders"
)

func InitTicTac(gameRoom *GameRoom) {
	gameBoard := games.InitBoard()
	
	for processed := range gameRoom.RuleChan {
		if processed.ID == messages.TTT_MOVE_ID {
			move := processed.MoveM.Move
			if games.IsValidMove(move, gameBoard) == 1 {
				gameBoard = games.MakeMove(move, processed.MoveM.Player, gameBoard)
				
				processed.MoveM.GameBoard = gameBoard
				processed.MoveM.HasWon = games.HasWon(gameBoard)
				processed.MoveM.IsDraw = games.IsDraw(gameBoard)
				
				if processed.MoveM.IsDraw == 1 || processed.MoveM.HasWon != 0 {
					gameBoard = games.ClearBoard(gameBoard)
				}
				processed.MoveM.IsValid = 1
			} else {
				processed.MoveM.IsValid = 0
			}
			gameRoom.SendMult <- MultipleMessage{encoders.EncodeMoveMessage(processed.MoveM), gameRoom.roomData.CS.ClientCount, gameRoom.roomData.CS.Clients}

		} else if processed.ID == messages.KICK_ID {
			gameBoard = games.ClearBoard(gameBoard)
			gameRoom.SendMult <- MultipleMessage{encoders.EncodeMoveMessage(processed.MoveM), gameRoom.roomData.CS.ClientCount, gameRoom.roomData.CS.Clients}
		} else if processed.ID == messages.ROOM_CLOSED_ID {
			break
		}
	}
}

