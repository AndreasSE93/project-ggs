package game

import(
	"server/messages"
	"server/games"
	"server/encoders"
)

func InitTicTac(gameRoom *GameRoom) {
	gameBoard := games.InitBoard()
	
	for {
		processed := <- gameRoom.RuleChan
		move := processed.MoveM.Move
		if games.IsValidMove(move, gameBoard) == 1 {
			gameBoard = games.MakeMove(move, processed.MoveM.Player, gameBoard)
			
			processed.MoveM.GameBoard = gameBoard
			processed.MoveM.HasWon = games.HasWon(gameBoard)
			processed.MoveM.IsDraw = games.IsDraw(gameBoard)

			if processed.MoveM.IsDraw == 1 || processed.MoveM.HasWon != 0 {
				games.ClearBoard(gameBoard)
			}
			processed.MoveM.IsValid = 1
		} else {
			processed.MoveM.IsValid = 0
		}
		go sendImmediateMessage(encoders.EncodeMoveMessage(messages.TTT_MOVE_ID, processed.MoveM), gameRoom.roomData.CS)
	}
}

