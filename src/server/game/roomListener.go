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
}

func sendToSingle(message string, conn connection.Connector) {
	conn.Connection.Write([]byte(message + "\n"))
}

func sendImmediateMessage(message string, cs messages.ClientSection) {
	fmt.Printf("Sending to %d clients: %s\n", cs.ClientCount, message)
	for i := 0; i < cs.ClientCount; i++ {
		cs.Clients[i].Connection.Write([]byte(message + "\n"))
	}
}

func gameRoomListener(gameRoom *GameRoom) {
	go sendImmediateMessage(encoders.EncodeHostedRoom(gameRoom.roomData),gameRoom.roomData.CS)
	go sendImmediateMessage(encoders.EncodeHostedRoom(gameRoom.roomData),gameRoom.roomData.CS)
	for {
		processed := <- gameRoom.roomData.SS.GameChan
		
		if processed.ID == messages.CHAT_ID {
			go sendImmediateMessage(encoders.EncodeChatMessage(processed.ChatM, processed.Origin), gameRoom.roomData.CS)

		} else if processed.ID == messages.JOIN_ID {
			gameRoom.roomData = gameRoom.lm.GetRoom(gameRoom.roomData.CS.RoomID)
			go sendImmediateMessage(encoders.EncodeJoinedRoom(&gameRoom.roomData), gameRoom.roomData.CS)
			go sendImmediateMessage(encoders.EncodeJoinedRoom(&gameRoom.roomData), gameRoom.roomData.CS)
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
			go sendImmediateMessage(encoders.EncodeJoinedRoom(&gameRoom.roomData), gameRoom.roomData.CS)
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
				}
			}
		}
		gameRoom.RuleChan <- processed
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
		go InitAchtung(game)
		go gameRoomListener(game)
		
	default: go gameRoomListener(game)
	}
}
