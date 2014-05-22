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

func messageInterpreter(messageTransfer chan string, sendToLobby chan messages.ProcessedMessage, client connection.Connector, updateClient chan connection.Connector) {
	defer close(sendToLobby)
	for {
		pMsg := new(messages.ProcessedMessage)
		message, ok := <- messageTransfer
		if !ok {
			return
		}
		json.Unmarshal([]byte(message), pMsg)

		select {
		case client = <-updateClient:
		default:
		}
		pMsg.Origin = client

		if pMsg.ID == messages.INIT_ID {
			initM := new(messages.InitMessage)
			json.Unmarshal([]byte(message), initM)
			pMsg.InitM = *initM
		} else if pMsg.ID == messages.CHAT_ID {
			chatM := new(messages.ChatMessage)
			json.Unmarshal([]byte(message), chatM)
			pMsg.ChatM = *chatM

		} else if pMsg.ID == messages.HOST_ID {
			hostM := new(messages.HostNew)
			json.Unmarshal([]byte(message), hostM)
			pMsg.Host = *hostM

		} else if pMsg.ID == messages.JOIN_ID {
			joinM := new(messages.JoinExisting)
			json.Unmarshal([]byte(message), joinM) 
			pMsg.Join = *joinM

		} else if pMsg.ID == messages.REFRESH_ID {
			updateM := new(messages.UpdateRooms)
			json.Unmarshal([]byte(message), updateM)
			pMsg.Update = *updateM

		} else if pMsg.ID == messages.TTT_CHAT_ID {
			chatM := new(messages.ChatMessage)
			json.Unmarshal([]byte(message), chatM)
			pMsg.ChatM = *chatM
			
		} else if pMsg.ID == messages.TTT_MOVE_ID {
			moveM := new(messages.MoveMessage)
			json.Unmarshal([]byte(message), moveM)
			pMsg.MoveM = *moveM
		} else if pMsg.ID == messages.START_ID {
			startM := new(messages.UpdateRooms)
			json.Unmarshal([]byte(message), startM)
			pMsg.Update = *startM
		} else if pMsg.ID == messages.SNAKES_CLIENT_ID {
			snakes := new(messages.SnakesEvent)
			json.Unmarshal([]byte(message), snakes)
			pMsg.Snakes = *snakes
		} else if pMsg.ID == messages.KICK_ID {
			kick := new(messages.KickMessage)
			json.Unmarshal([]byte(message), kick)
			pMsg.Kick = *kick
		} else {
			fmt.Printf("Unknown packet received: %+v\n", *pMsg)
			continue
		}
		sendToLobby <- *pMsg
	}
}

func ActivateReceiver(messageProcessing chan messages.ProcessedMessage, client connection.Connector, updateClient chan connection.Connector) {
	messageTransfer := make(chan string)
	defer close(messageTransfer)
	go messageInterpreter(messageTransfer, messageProcessing, client, updateClient)

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

	updateClient := make(chan connection.Connector)
	
	go ActivateReceiver(processedChan, core.client, updateClient)
	go ActivateSender(finJSON, core.client)

	refreshList := ReqUpdate(messages.UpdateRooms{}, *core)
	finJSON <- encoders.EncodeRefreshList(refreshList)

	processed := <-processedChan
	client = db.Get(interface{}(client.ConnectorID)).(connection.Connector)
	client.UserName = processed.InitM.UserName
	db.Add(interface{}(client.ConnectorID), interface{}(client))
	updateClient <-client
	fmt.Println("after first processed")
	cr.Connect(client)

	lobbyChan := make(chan messages.ProcessedMessage)
	gameChan := lobbyChan
	go receiveLobbyChat(lobbyChan, cr)

	defer lm.Kick(client)

	for {
		processed, ok := <- processedChan
		if !ok {
			return
		}
		//fmt.Printf("Received decoded message: %v\n", processed)

		if processed.ID == messages.CHAT_ID {
			gameChan <- processed
	
		} else if processed.ID == messages.HOST_ID {
			hostedRoom, hostedErr := ReqHost(processed.Host, *core)
			if hostedErr == nil {
				gameChan = hostedRoom.SS.GameChan
				cr.Disconnect(processed.Origin)
				go game.CreateGameRoom(hostedRoom, core.lm)
			} else {
				fmt.Printf("Unable to host room: %s\n", hostedErr)
			}

		} else if processed.ID == messages.JOIN_ID {
			joinedRoom := ReqJoin(processed.Join, *core)
			fmt.Println(joinedRoom.SS.GameChan)
			gameChan = joinedRoom.SS.GameChan
			gameChan <- processed
			cr.Disconnect(processed.Origin)

		} else if processed.ID == messages.REFRESH_ID {
			refreshList := ReqUpdate(processed.Update, *core)
			finJSON <- encoders.EncodeRefreshList(refreshList)

		} else if processed.ID == messages.TTT_CHAT_ID {
			gameChan <- processed

		} else if processed.ID == messages.TTT_MOVE_ID {
			gameChan <- processed
			
		} else if processed.ID == messages.START_ID {
			gameChan <- processed

		} else if processed.ID == messages.KICK_ID {
			gameChan <- processed
			gameChan = lobbyChan
			cr.Connect(processed.Origin)

		} else if processed.ID == messages.SNAKES_CLIENT_ID {
			gameChan <-processed
		} else {
			fmt.Println("Something went wrong!")
		}
		//fmt.Printf("lm=%+v\n", lm.GetShadow())
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
