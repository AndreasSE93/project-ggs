package games

func InitBoard() []int {
	gameBoard := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	return gameBoard
}

func ClearBoard(gameBoard []int) []int {
	gameBoard = InitBoard()
	return gameBoard
}

func IsValidMove(move int, gameBoard []int) int {
	if gameBoard[move] == 0 {
		return 1
	} else {
		return 0
	}
}

func HasWon(field []int) int {

	// Horizontal
	if field[0] == field[1] && field[0] == field[2] && field[0] != 0 {
		return field[0]
	}
	if field[3] == field[4] && field[3] == field[5] && field[3] != 0 {
		return field[3]
	}
	if field[6] == field[7] && field[6] == field[8] && field[6] != 0 {
		return field[6]
	}

	// Vertical
	if field[0] == field[3] && field[0] == field[6] && field[0] != 0 {
		return field[0]
	}
	if field[1] == field[4] && field[1] == field[7] && field[1] != 0 {
		return field[4]
	}
	if field[2] == field[5] && field[2] == field[8] && field[2] != 0 {
		return field[2]
	}

	// Diagonal
	if field[0] == field[4] && field[0] == field[8] && field[0] != 0 {
		return field[0]
	}
	if field[2] == field[4] && field[2] == field[6] && field[2] != 0 {
		return field[2]
	} else {
		return 0
	}

}

func IsDraw(gameBoard []int) int {
	for i := 0; i < 9; i++ {
		if gameBoard[i] == 0 {
			return 0
		}

	}
	return 1

}

func MakeMove(move int, player int, gameBoard []int) []int {
	gameBoard[move] = player
	return gameBoard
}
