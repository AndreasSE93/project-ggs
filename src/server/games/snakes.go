package games

import (
	//"fmt"
	"math"
	"time"
	"math/rand"
	"server/messages"
)

const (
	GAME_W = 835
	GAME_H = 790
	TURN_DELTA = 5
)

type achtungPixel struct {
	Player byte
	Creation uint16
}

type achtungBoard *[GAME_W][GAME_H]achtungPixel 

func InitBoardSnakes() achtungBoard {
	return new([GAME_W][GAME_H]achtungPixel)
}

func InitAchtungPlayerArray(numPlayers int) (pA []messages.Player) {
	pA = make([]messages.Player, numPlayers)
	for player := range(pA) {
		pA[player] = InitAchtungPlayer(pA[player])
		pA[player].PlayerID = player + 1
	}
	return
}

func ResetAchtungPlayerArray(pA []messages.Player) {
	for player := range(pA) {
		pA[player] = InitAchtungPlayer(pA[player])
	}
	return
}

func InitAchtungPlayer(player messages.Player) messages.Player {
	rand.Seed(time.Now().UTC().UnixNano())
	player.PosX = rand.Float64() * (700 - 135) + 135
	player.PosY = rand.Float64() * (655 - 135) + 135
	player.Alive = true
	player.Direction = rand.Float64() * 360
	return player
}

func addPoints(playerArray []messages.Player, points int) {
	for i := range playerArray {
		if playerArray[i].Alive {
			playerArray[i].Score += points
		}
	}
}

func LegalMoveSnakes (player *messages.Player, gameBoard achtungBoard, playerArray []messages.Player, time uint16) bool {
	if !player.Alive {
		return false
	}

	if player.PosX >= GAME_W - 1 || player.PosX < 1 || player.PosY >= GAME_H - 1 || player.PosY < 1 {
		player.Alive = false
		addPoints(playerArray, 1)
		return false
	}
	
	if id := gameBoard[int(player.PosX)][int(player.PosY)]; id.Player != 0 && (time - id.Creation > 360 / TURN_DELTA / 2  || id.Player != byte(player.PlayerID))  {
		println("Player", player.PlayerID, "hit player", id.Player, id.Creation, time, 360 / TURN_DELTA / 2)
		player.Alive = false
		addPoints(playerArray, 1)
		return false
	}
	return true
}

func UpdateMoveSnakes (player messages.Player, move string) messages.Player{
	if move == "VK_LEFT"{
		player.Direction -= TURN_DELTA
	}
	if move == "VK_RIGHT"{
		player.Direction += TURN_DELTA
	}
	return player
}

func UpdateAllMovesSnakes(playerArray []messages.Player, gameBoard achtungBoard, time uint16) []messages.Player{
	for i:=0; i<4; i++{
		if playerArray[i].Alive {
			playerArray[i].Direction = math.Mod(playerArray[i].Direction, 360)
			playerArray[i].PosX += 1 * math.Cos(playerArray[i].Direction*(math.Pi/180.0))
			playerArray[i].PosY += 1 * math.Sin(playerArray[i].Direction*(math.Pi/180.0))	

			if LegalMoveSnakes(&playerArray[i], gameBoard, playerArray, time) {
				playerArray[i].PosArray[0] = playerArray[i].PosX
				playerArray[i].PosArray[1] = playerArray[i].PosY

				playerArray[i].PosArray[2] = playerArray[i].PosX+1
				playerArray[i].PosArray[3] = playerArray[i].PosY
				
				playerArray[i].PosArray[4] = playerArray[i].PosX-1
				playerArray[i].PosArray[5] = playerArray[i].PosY
				
				playerArray[i].PosArray[6] = playerArray[i].PosX
				playerArray[i].PosArray[7] = playerArray[i].PosY+1
				
				playerArray[i].PosArray[8] = playerArray[i].PosX
				playerArray[i].PosArray[9] = playerArray[i].PosY -1
			}
		}
	}

	return playerArray
}

func AchtungFinished(playerArray []messages.Player, gameBoard achtungBoard) bool {
	alivePlayers := 0
	for _, player := range(playerArray) {
		if player.Alive {
			alivePlayers++
		}
	}
	if alivePlayers <= 1 {
		ResetAchtungPlayerArray(playerArray)
		*gameBoard = *InitBoardSnakes()
		return true
	} else {
		return false
	}
}

func fillPixel(gameBoard achtungBoard, x, y int, n byte, time uint16) {
	if x >=0 && x < GAME_W && y >=0 && y < GAME_H  {
		if gameBoard[x][y].Player == 0{
			gameBoard[x][y] = achtungPixel{
				Player:   n,
				Creation: time,
			}
		}
	}
}



/*

X X X
X O X
X X X


*/



func DoMove (playerArray []messages.Player, gameBoard achtungBoard, time uint16) {
	for i := 0; i<4; i++ {
		if LegalMoveSnakes(&playerArray[i], gameBoard, playerArray, time)  {
			fillPixel(gameBoard, int(playerArray[i].PosX  ), int(playerArray[i].PosY  ), byte(i + 1), time)

			fillPixel(gameBoard, int(playerArray[i].PosX  ), int(playerArray[i].PosY+1), byte(i + 1), time)
			fillPixel(gameBoard, int(playerArray[i].PosX  ), int(playerArray[i].PosY-1), byte(i + 1), time)
			fillPixel(gameBoard, int(playerArray[i].PosX+1), int(playerArray[i].PosY  ), byte(i + 1), time)
			fillPixel(gameBoard, int(playerArray[i].PosX-1), int(playerArray[i].PosY  ), byte(i + 1), time)


			fillPixel(gameBoard, int(playerArray[i].PosX+1  ), int(playerArray[i].PosY+1), byte(i + 1), time)
			fillPixel(gameBoard, int(playerArray[i].PosX-1  ), int(playerArray[i].PosY-1), byte(i + 1), time)
			fillPixel(gameBoard, int(playerArray[i].PosX+1), int(playerArray[i].PosY-1  ), byte(i + 1), time)
			fillPixel(gameBoard, int(playerArray[i].PosX-1), int(playerArray[i].PosY+1  ), byte(i + 1), time)

		}
	}
}

