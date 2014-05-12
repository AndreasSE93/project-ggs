package game

import(
//	"fmt"
//	"encoding/json"
	"server/database/lobbyMap"
	"server/messages"
	"server/encoders"
	"server/connection"
)

type GameRoom struct {
	hr messages.HostRoom
	lm *lobbyMap.LobbyMap
}

func sendImmediateMessage(message string, clients []connection.Connector) {
	for c := range clients {
		clients[c].Connection.Write([]byte(message + "\n"))
	}
}

func gameRoomListener(gr *GameRoom) {
	for {
		processed := <- gr.hr.GameChan
		
		if processed.ID == messages.CHAT_ID {
			go sendImmediateMessage(encoders.EncodeChatMessage(messages.CHAT_ID, processed.ChatM), gr.hr.Clients)
		} else if processed.ID == messages.JOIN_ID {
			gr.hr = gr.lm.GetRoom(gr.hr.RoomID)
			go sendImmediateMessage(encoders.EncodeJoinedRoom(messages.JOIN_ID, &gr.hr), gr.hr.Clients)
		}
	}
}


func CreateGameRoom(hr messages.HostRoom, lm *lobbyMap.LobbyMap) {
	game := new(GameRoom)
	game.hr = hr
	game.lm = lm

	go gameRoomListener(game)
}
