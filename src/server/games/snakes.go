package games


import (
	
//	"fmt"
	"math"
	"time"
	"math/rand"
	"server/messages"
)



func InitBoardSnakes() [835][790]float64 {
	var g [835][790]float64
	return g
}

func LegalMoveSnakes (player messages.Player, gameBoard [835][790]float64) bool{
	if player.PosX >= 834 || player.PosX < 1 ||player.PosY >= 789 || player.PosY < 1 {
		player.Alive = false
	//	fmt.Println("LeGALMOVE            1INSIDE GMAEBOARD ", player.PosX, " ", player.PosY)
		return false
	}
	
	/*if gameBoard[int(player.PosX)][int(player.PosY)] == 1 {
		player.Alive = false
		fmt.Println(gameBoard[int(player.PosX)][int(player.PosY)], "INSIDE GMAEBOARD ", player.PosX, " ", player.PosY )
		return false
	}*/
	//fmt.Println("LeGALMOVE ", player.PosX," " ,player.PosY )
	return true
}

func UpdateMoveSnakes (player messages.Player, move string) messages.Player{
	
	if move == "VK_LEFT"{
		player.Direction -= 10
	}
	if move == "VK_RIGHT"{
		player.Direction += 10
	}

	
	return player
}

func MakePlayerSnakes(ID int) messages.Player{
	rand.Seed(time.Now().UTC().UnixNano())
	var x [10]float64
	player := messages.Player{ "",ID, x, float64(rand.Intn(700-135)+135), float64(rand.Intn(655-135)+135), true, 0, float64(rand.Intn(360))}
	return player
}


func UpdateAllMovesSnakes(playerArray []messages.Player, gameBoard [835][790]float64 ) []messages.Player{
	
	for i:=0; i<4; i++{
		if LegalMoveSnakes (playerArray[i], gameBoard) {
			playerArray[i].Direction = math.Mod(playerArray[i].Direction, 360)
			playerArray[i].PosX += 2 * math.Cos(playerArray[i].Direction*(math.Pi/180.0))
			playerArray[i].PosY += 2 * math.Sin(playerArray[i].Direction*(math.Pi/180.0))	
		
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
	

	return playerArray
}


func DoMove (playerArray []messages.Player, gameBoard [835][790]float64) [835][790]float64 {
	
	for i:= 0 ; i<4; i++{
		if LegalMoveSnakes(playerArray[i], gameBoard)  {
			gameBoard[int(playerArray[i].PosX)][int(playerArray[i].PosY)] = 1
			gameBoard[int(playerArray[i].PosX)][int(playerArray[i].PosY+1)] = 1
			gameBoard[int(playerArray[i].PosX)][int(playerArray[i].PosY-1)] = 1
			gameBoard[int(playerArray[i].PosX+1)][int(playerArray[i].PosY)] = 1
			gameBoard[int(playerArray[i].PosX-1)][int(playerArray[i].PosY)] = 1

		}
	}
	return gameBoard
}

