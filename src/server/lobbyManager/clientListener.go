package lobbyManager

import (
	"fmt"
	"encoding/json"
	"server/connection"
	"server/database"
	"server/database/lobbyMap"
	"server/encoders"
	"server/messages"
	"server/chatRoom"
	"server/game"
)

type ClientCore struct {
	client connection.Connector
	lm *lobbyMap.LobbyMap
}

func messageInterpreter(messageTransfer chan string, sendToLobby chan messages.ProcessedMessage, client connection.Connector) {
	defer close(sendToLobby)
	for message := range messageTransfer {
		pMsg := new(messages.ProcessedMessage)
		json.Unmarshal([]byte(message), pMsg)
		pMsg.Origin = client

		switch pMsg.ID {
		case messages.INIT_ID:
			initM := new(messages.InitMessage)
			json.Unmarshal([]byte(message), initM)
			pMsg.InitM = *initM
			client.UserName = initM.UserName
		case messages.CHAT_ID:
			chatM := new(messages.ChatMessage)
			json.Unmarshal([]byte(message), chatM)
			pMsg.ChatM = *chatM
		case messages.HOST_ID:
			hostM := new(messages.HostNew)
			json.Unmarshal([]byte(message), hostM)
			pMsg.Host = *hostM
		case messages.JOIN_ID:
			joinM := new(messages.JoinExisting)
			json.Unmarshal([]byte(message), joinM) 
			pMsg.Join = *joinM
		case messages.REFRESH_ID:
			updateM := new(messages.UpdateRooms)
			json.Unmarshal([]byte(message), updateM)
			pMsg.Update = *updateM
		case messages.TTT_CHAT_ID:
			chatM := new(messages.ChatMessage)
			json.Unmarshal([]byte(message), chatM)
			pMsg.ChatM = *chatM
		case messages.TTT_MOVE_ID:
			moveM := new(messages.MoveMessage)
			json.Unmarshal([]byte(message), moveM)
			pMsg.MoveM = *moveM
		case messages.START_ID:
			startM := new(messages.UpdateRooms)
			json.Unmarshal([]byte(message), startM)
			pMsg.Update = *startM
		case messages.SNAKES_CLIENT_ID:
			snakes := new(messages.SnakesEvent)
			json.Unmarshal([]byte(message), snakes)
			pMsg.Snakes = *snakes
		case messages.KICK_ID:
			kick := new(messages.KickMessage)
			json.Unmarshal([]byte(message), kick)
			pMsg.Kick = *kick
		default:
			fmt.Printf("Unknown packet received: %+v\n", *pMsg)
			continue
		}
		sendToLobby <- *pMsg
	}
}

func ActivateReceiver(messageProcessing chan messages.ProcessedMessage, client connection.Connector) {
	messageTransfer := make(chan string)
	defer close(messageTransfer)
	go messageInterpreter(messageTransfer, messageProcessing, client)

	for client.Scanner.Scan() {
		//fmt.Printf("Received message from client %d: %+v\n", client.ConnectorID, client.Scanner.Text())
		messageTransfer <- client.Scanner.Text()
	}
	if err := client.Scanner.Err(); err != nil {
		fmt.Printf("Error reading from client %d: %s\n", client.ConnectorID, err)
	} else {
		fmt.Printf("Client %d disconnected\n", client.ConnectorID)
	}
}

func ActivateSender(serverTerminal chan string, client connection.Connector) {
	for {
		jsonString := <- serverTerminal
		//fmt.Printf("Sending to client %d: %s\n", client.ConnectorID, jsonString)
		client.Connection.Write([]byte(jsonString + "\n"))
	}
}

func ClientListener(lm *lobbyMap.LobbyMap, db *database.Database, client connection.Connector, cr *chatRoom.ChatRoom) {
	defer client.Connection.Close()

	core := new(ClientCore)
	core.client = client
	core.lm = lm
	processedChan := make(chan messages.ProcessedMessage)

	finJSON := make(chan string)
	
	go ActivateReceiver(processedChan, core.client)
	go ActivateSender(finJSON, core.client)

	refreshList := ReqUpdate(messages.UpdateRooms{}, *core)
	finJSON <- encoders.EncodeRefreshList(refreshList)

	processed := <-processedChan
	client = db.Get(interface{}(client.ConnectorID)).(connection.Connector)
	client.UserName = processed.InitM.UserName
	db.Add(interface{}(client.ConnectorID), interface{}(client))
	cr.Connect(client)

	lobbyChan := make(chan messages.ProcessedMessage)
	gameChan := lobbyChan
	go receiveLobbyChat(lobbyChan, cr)

	defer lm.Kick(client)

	for processed := range processedChan {
		switch processed.ID {
		case messages.CHAT_ID:
			gameChan <- processed
		case messages.HOST_ID:
			hostedRoom, hostedErr := ReqHost(processed.Host, *core)
			if hostedErr == nil {
				gameChan = hostedRoom.SS.GameChan
				cr.Disconnect(processed.Origin)
				go game.CreateGameRoom(hostedRoom, core.lm)
			} else {
				fmt.Printf("Unable to host room: %s\n", hostedErr)
			}
		case messages.JOIN_ID:
			joinedRoom := ReqJoin(processed.Join, *core)
			fmt.Println(joinedRoom.SS.GameChan)
			gameChan = joinedRoom.SS.GameChan
			gameChan <- processed
			cr.Disconnect(processed.Origin)
		case messages.REFRESH_ID:
			refreshList := ReqUpdate(processed.Update, *core)
			finJSON <- encoders.EncodeRefreshList(refreshList)
		case messages.TTT_CHAT_ID:
			gameChan <- processed
		case messages.TTT_MOVE_ID:
			gameChan <- processed
		case messages.START_ID:
			gameChan <- processed
		case messages.KICK_ID:
			gameChan <- processed
			lm.Kick(client)
			gameChan = lobbyChan
			cr.Connect(processed.Origin)
			refreshList := ReqUpdate(messages.UpdateRooms{}, *core)
			finJSON <- encoders.EncodeRefreshList(refreshList)

		case messages.SNAKES_CLIENT_ID:
			gameChan <-processed
		default:
			fmt.Println("Something went wrong!")
		}
	}
}

func receiveLobbyChat(ch chan messages.ProcessedMessage, cr *chatRoom.ChatRoom) {
	for msg := range ch {
		if msg.ID == messages.CHAT_ID {
			cr.SendMessage(msg)
		} else {
			fmt.Println("ATE RANDOM PACKAGE", msg.ID)
		}
	}
}
