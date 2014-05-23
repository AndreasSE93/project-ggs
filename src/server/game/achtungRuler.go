package game

import(
	"time"
	"server/encoders"
	"server/games"
	"server/messages"
	"fmt"
	"math/rand"
)

func snakeListener(gameRoom *GameRoom) {
	for {
		processed, ok := <- gameRoom.RuleChan
		if !ok {
			return
		} else if processed.ID == messages.START_ID {
			break
		}
	}
	newGameRoom := make(chan messages.RoomData)
	termChan    := make(chan interface{})
	fmt.Println(gameRoom.roomData.CS.ClientCount)
	PlayerArray := games.InitAchtungPlayerArray(gameRoom.roomData.CS.ClientCount, gameRoom.roomData.CS.MaxSize)
	setNames(gameRoom, PlayerArray)
	go snakesHandler(PlayerArray, gameRoom, newGameRoom, termChan)
	for processed := range gameRoom.RuleChan {
		switch processed.ID {
		case messages.SNAKES_CLIENT_ID:
			fmt.Println("ID",processed.Snakes.PlayerID)
			PlayerArray[processed.Snakes.PlayerID-1] = games.UpdateMoveSnakes(PlayerArray[processed.Snakes.PlayerID-1], processed.Snakes.Move)
		case messages.JOIN_ID:
			newGameRoom <- gameRoom.lm.GetRoom(gameRoom.roomData.CS.RoomID)
		case messages.ROOM_CLOSED_ID:
			termChan <- nil
		case messages.KICK_ID:
			for player:= range PlayerArray {
				if PlayerArray[player].PlayerName == processed.Origin.UserName {
					PlayerArray[player] = messages.Player{}
				}
			}
		}
	}
}


func setNames (gameRoom *GameRoom, pA []messages.Player){
	for i := 0; i < len(pA); i++ {
		pA[i].PlayerName = gameRoom.roomData.CS.Clients[i].UserName
	}

}

func findSmallestID(pA []messages.Player) int {

	for player:= range pA {
		if pA[player].PlayerID == 0 {
			return player
		}
	}

return -1
}


func snakesHandler(pA []messages.Player, gameRoom *GameRoom, newGameRoom chan messages.RoomData, termChan chan interface{}) { 
	gameBoard := games.InitBoardSnakes()
	
//	sMess := SingleMessage{}
	msg := MultipleMessage{}
	msg.ClientCount = gameRoom.roomData.CS.ClientCount
	skip := true

	var Time uint16
	for {
		select {
		case newerGameRoom := <- newGameRoom:
			gameRoom.roomData = newerGameRoom
			msg.ClientCount = gameRoom.roomData.CS.ClientCount
		case <- termChan:
			return
	
		default:
		}
		
		if Time % 15 == 0 {
			rand.Seed(time.Now().UTC().UnixNano())
			x := rand.Intn(100)
			if x < 3 {
				skip = false
			} else {
				skip = true
			}
			
		}
	
		pA = games.UpdateAllMovesSnakes(pA, gameBoard, Time, skip)
		games.DoMove(pA, gameBoard, Time, skip)
		clear := games.AchtungFinished(pA, gameBoard)
	
		msg.message = encoders.EncodeSnakeMessage(messages.SNAKES_MOVES_ID, pA, clear, false, "")
		msg.conn = gameRoom.roomData.CS.Clients
		gameRoom.SendMult <- msg
		
		if clear {
			/*hasWon, WinnerName := games.WinnerWinnerChickenDinner(pA)
			if hasWon {
				msg.message = encoders.EncodeSnakeMessage(messages.SNAKES_MOVES_ID, pA, clear, hasWon, WinnerName)
				msg.conn = gameRoom.roomData.CS.Clients
				gameRoom.SendMult <- msg
				pA =  games.InitAchtungPlayerArray(gameRoom.roomData.CS.ClientCount, gameRoom.roomData.CS.MaxSize)
				setNames(gameRoom, pA)
				
				
				
			}*/
			Time = 0
			time.Sleep(3 * time.Second)
		
			
			
		} else {
			time.Sleep(15 * time.Millisecond )
		}
		

		Time++
	}
}


