package game

import(
	"fmt"
	"server/database/lobbyMap"
	"server/messages"
	"server/encoders"
	"server/connection"
)

type GameRoom struct {
	roomData messages.RoomData
	lm *lobbyMap.LobbyMap
	GameType string
	RuleChan chan messages.ProcessedMessage
	Started, Startable bool
	SendSingle chan SingleMessage
	SendMult chan MultipleMessage
}

type SingleMessage struct {
	message string
	conn connection.Connector
}

type MultipleMessage struct {
	message string
	ClientCount int
	conn []connection.Connector
}

func sendToSingle(SendOutChan chan SingleMessage) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Failed to send to single player")
			return
		}
	}()

	for {
		message := <- SendOutChan
		message.conn.Connection.Write([]byte(message.message + "\n"))
	}
}

func sendImmediateMessage(SendOutChan chan MultipleMessage) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Failed to send to single player")
			return
		}
	}()

	for {
		fmt.Println(len(SendOutChan))
		message := <- SendOutChan
//		fmt.Printf("Sending to %d clients: %s\n", message.ClientCount, message.message)
		for i := 0; i < message.ClientCount; i++ {
			message.conn[i].Connection.Write([]byte(message.message + "\n"))
		}
	}
}

func gameRoomListener(gameRoom *GameRoom) {
	defer fmt.Printf("Room %d closed\n", gameRoom.roomData.CS.RoomID)
	gameRoom.SendMult <- MultipleMessage{encoders.EncodeHostedRoom(gameRoom.roomData), gameRoom.roomData.CS.ClientCount, gameRoom.roomData.CS.Clients}
	for {
		processed := <- gameRoom.roomData.SS.GameChan
		fmt.Println("INSIDE ROOM",processed)
		
		if processed.ID == messages.CHAT_ID {
			gameRoom.SendMult <- MultipleMessage{encoders.EncodeChatMessage(processed.ChatM, processed.Origin), gameRoom.roomData.CS.ClientCount, gameRoom.roomData.CS.Clients}

		} else if processed.ID == messages.JOIN_ID {
			if room := gameRoom.lm.GetRoom(gameRoom.roomData.CS.RoomID); room.CS.RoomID > 0 {
				gameRoom.roomData = room
			} else {
				break
			}
			gameRoom.SendMult <- MultipleMessage{encoders.EncodeJoinedRoom(gameRoom.roomData), gameRoom.roomData.CS.ClientCount, gameRoom.roomData.CS.Clients}
			if gameRoom.roomData.CS.ClientCount == gameRoom.roomData.CS.MaxSize {
				gameRoom.Startable = true
				gameRoom.SendSingle <- SingleMessage{encoders.EncodeStartable(true), gameRoom.roomData.CS.Clients[0]}
			}

		} else if processed.ID == messages.KICK_ID {
			if room := gameRoom.lm.Kick(processed.Origin); room != nil {
				gameRoom.roomData = *room
			} else {
				break
			}
			gameRoom.SendSingle <- SingleMessage{encoders.EncodeKick(), processed.Origin}
			gameRoom.SendSingle <- SingleMessage{encoders.EncodeStartable(false), gameRoom.roomData.CS.Clients[0]}
			gameRoom.SendMult <- MultipleMessage{encoders.EncodeJoinedRoom(gameRoom.roomData), gameRoom.roomData.CS.ClientCount, gameRoom.roomData.CS.Clients}
			if gameRoom.Started == true {
				gameRoom.Started = false
				gameRoom.Startable = false
				gameRoom.SendMult <- MultipleMessage{encoders.EncodeStartGame(false, 0), gameRoom.roomData.CS.ClientCount, gameRoom.roomData.CS.Clients}
			}

		} else if processed.ID == messages.START_ID {
			if gameRoom.Startable {
				gameRoom.Started = true
				fmt.Printf("Game %d started\n", gameRoom.roomData.CS.RoomID)
				for place := 0; place < gameRoom.roomData.CS.ClientCount; place++ {
					gameRoom.SendSingle <- SingleMessage{encoders.EncodeStartGame(true, place+1), gameRoom.roomData.CS.Clients[place]}
				}
			}
		}
		gameRoom.RuleChan <- processed
	}
	gameRoom.RuleChan <- messages.ProcessedMessage{
		ID: messages.ROOM_CLOSED_ID,
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
	game.SendSingle = make(chan SingleMessage)
	game.SendMult = make(chan MultipleMessage, 10)

	go sendToSingle(game.SendSingle)
	go sendImmediateMessage(game.SendMult)
	
	switch game.GameType {
	case "TicTacToe": 
		go InitTicTac(game)
		go gameRoomListener(game)

	case "Achtung Die Kurve":
		go snakeListener(game)
		go gameRoomListener(game)
		
	default: go gameRoomListener(game)
	}
}
