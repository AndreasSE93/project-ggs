package game

import(
	"fmt"
	"server/database/lobbyMap"
	"server/messages"
	"server/encoders"
	"server/connection"
	"time"
	"server/games"
)

type GameRoom struct {
	roomData messages.RoomData
	lm *lobbyMap.LobbyMap
	GameType string
	RuleChan chan messages.ProcessedMessage
	Started, Startable bool
}

func sendToSingle(message string, conn connection.Connector) {
	conn.Connection.Write([]byte(message + "\n"))
}

func sendImmediateMessage(message string, cs messages.ClientSection) {
	//fmt.Printf("Sending to %d clients: %s\n", cs.ClientCount, message)
	for i := 0; i < cs.ClientCount; i++ {
		cs.Clients[i].Connection.Write([]byte(message + "\n"))
		//fmt.Println(cs.ClientCount)
	}
}

func gameRoomListener(gameRoom *GameRoom) {
	go sendImmediateMessage(encoders.EncodeHostedRoom(gameRoom.roomData),gameRoom.roomData.CS)
	for {
		processed := <- gameRoom.roomData.SS.GameChan
		
		if processed.ID == messages.CHAT_ID {
			go sendImmediateMessage(encoders.EncodeChatMessage(processed.ChatM, processed.Origin), gameRoom.roomData.CS)

		} else if processed.ID == messages.JOIN_ID {
			gameRoom.roomData = gameRoom.lm.GetRoom(gameRoom.roomData.CS.RoomID)
			go sendImmediateMessage(encoders.EncodeJoinedRoom(gameRoom.roomData), gameRoom.roomData.CS)
			if gameRoom.roomData.CS.ClientCount == gameRoom.roomData.CS.MaxSize {
				gameRoom.Startable = true
				go sendToSingle(encoders.EncodeStartable(true), gameRoom.roomData.CS.Clients[0])
			}

		} else if processed.ID == messages.KICK_ID {
			if room := gameRoom.lm.Kick(processed.Origin); room == nil {
				return
			} else {
				gameRoom.roomData = *room
			}
			go sendToSingle(encoders.EncodeKick(), processed.Origin)
			go sendToSingle(encoders.EncodeStartable(false), gameRoom.roomData.CS.Clients[0])
			go sendImmediateMessage(encoders.EncodeJoinedRoom(gameRoom.roomData), gameRoom.roomData.CS)
			if gameRoom.Started == true {
				gameRoom.Started = false
				gameRoom.Startable = false
				go sendImmediateMessage(encoders.EncodeStartGame(false, 0), gameRoom.roomData.CS)
			}

		} else if processed.ID == messages.START_ID {
			if gameRoom.Startable {
				gameRoom.Started = true
				for place := 0; place < gameRoom.roomData.CS.ClientCount; place++ {
					go sendToSingle(encoders.EncodeStartGame(true, place+1), gameRoom.roomData.CS.Clients[place])
					fmt.Println("SENDING STARTEDMESSAGE")
				}
			}
		}
		gameRoom.RuleChan <- processed
	}
}

// SNAKE FUNCTIONS PORTED FROM CLIENT BRANCH
func snakeListener(gameRoom *GameRoom){
	
	newGameRoom := make(chan messages.RoomData)
	PlayerArray := []messages.Player{games.MakePlayerSnakes(1),games.MakePlayerSnakes(2),games.MakePlayerSnakes(3),games.MakePlayerSnakes(4)}
	go snakesHandler(&PlayerArray, gameRoom, newGameRoom)
	for {
		processed := <- gameRoom.roomData.SS.GameChan
		if processed.ID == messages.SNAKES_CLIENT_ID {
			PlayerArray[processed.Snakes.PlayerID-1] = games.UpdateMoveSnakes(PlayerArray[processed.Snakes.PlayerID-1], processed.Snakes.Move)
		}

		if processed.ID == messages.JOIN_ID {
			newGameRoom <- gameRoom.lm.GetRoom(gameRoom.roomData.CS.RoomID)
		}

	}

}

func snakesHandler(pA *[]messages.Player, gameRoom *GameRoom, newGameRoom chan messages.RoomData) { 
	gameBoard := games.InitBoardSnakes()
	
	for{
		select{
		case newerGameRoom := <- newGameRoom:
			gameRoom.roomData = newerGameRoom
			
		default:

		}

		//gameRoom.roomData = gameRoom.lm.GetRoom(gameRoom.roomData.CS.RoomID)
		*pA = games.UpdateAllMovesSnakes( *pA, gameBoard)
		gameBoard = games.DoMove(*pA, gameBoard)
		go sendImmediateMessage(encoders.EncodeSnakeMessage(messages.SNAKES_MOVES_ID, *pA), gameRoom.roomData.CS)
		time.Sleep(25 * time.Millisecond )
	}
	

}

func CreateGameRoom(rd messages.RoomData, lm *lobbyMap.LobbyMap) {

	game := new(GameRoom)
	game.roomData = rd
	game.lm = lm
	game.RuleChan = make(chan messages.ProcessedMessage)
	game.GameType = rd.CS.GameName
	game.Started = false
	game.Startable = false
	
	switch game.GameType {
	case "TicTacToe": 
		go InitTicTac(game)
		go gameRoomListener(game)

	case "Achtung":
		go snakeListener(game)
		go gameRoomListener(game)
		
	default: go gameRoomListener(game)
	}
}
