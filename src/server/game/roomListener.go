package game

import(
	"fmt"
	"errors"
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
		message, ok := <- SendOutChan
		if ok {
			message.conn.Connection.Write([]byte(message.message + "\n"))
		} else {
			return
		}
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
		message, ok := <- SendOutChan
		if ok {
			for i := 0; i < message.ClientCount; i++ {
				message.conn[i].Connection.Write([]byte(message.message + "\n"))
			}
		} else {
			return
		}
	}
}

func gameRoomListener(gameRoom *GameRoom) {
	gameRoom.SendMult <- MultipleMessage{encoders.EncodeHostedRoom(gameRoom.roomData), gameRoom.roomData.CS.ClientCount, gameRoom.roomData.CS.Clients}
	for processed := range gameRoom.roomData.SS.GameChan {
		if processed.ID == messages.CHAT_ID {
			gameRoom.SendMult <- MultipleMessage{encoders.EncodeChatMessage(processed.ChatM, processed.Origin), gameRoom.roomData.CS.ClientCount, gameRoom.roomData.CS.Clients}

		} else if processed.ID == messages.JOIN_ID {
			if room := gameRoom.lm.GetRoom(gameRoom.roomData.CS.RoomID); room.CS.RoomID > 0 {
				gameRoom.roomData = room
			} else {
				break
			}
			gameRoom.SendMult <- MultipleMessage{encoders.EncodeJoinedRoom(gameRoom.roomData), gameRoom.roomData.CS.ClientCount, gameRoom.roomData.CS.Clients}
			if gameRoom.roomData.CS.ClientCount >= gameRoom.roomData.CS.MinSize {
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
	defer func() {
		lm.Delete(rd.CS.RoomID)
		if err := recover(); err == nil {
			defer fmt.Printf("Room %d closed\n", rd.CS.RoomID)
		} else {
			fmt.Printf("Room %d crashed: %s\n", rd.CS.RoomID, err)
			return
		}
	}()

	game := new(GameRoom)
	game.roomData = rd
	game.lm = lm
	game.RuleChan = make(chan messages.ProcessedMessage)
	defer close(game.RuleChan)
	game.GameType = rd.CS.GameName
	game.Started = false
	game.Startable = false
	game.SendSingle = make(chan SingleMessage)
	defer close(game.SendSingle)
	game.SendMult = make(chan MultipleMessage, 10)
	defer close(game.SendMult)

	go sendToSingle(game.SendSingle)
	go sendImmediateMessage(game.SendMult)

	switch game.GameType {
	case "TicTacToe":
		go InitTicTac(game)
	case "Achtung Die Kurve":
		go snakeListener(game)
	default:
		panic(errors.New("Invalid game type: " + game.GameType))
	}

	gameRoomListener(game)
}
