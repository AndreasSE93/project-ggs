package lobbyManager

import (
	"fmt"
	"server/connection"
	"server/database/lobbyMap"
	"server/encoders"
	"server/messages"
	"encoding/json"
	"server/game"
)

type ClientCore struct {
	client connection.Connector
	lm *lobbyMap.LobbyMap
}

func messageInterpreter(messageTransfer chan string, sendToLobby chan messages.ProcessedMessage) {
	for {
		pMsg := new(messages.ProcessedMessage)
		message := <- messageTransfer
		json.Unmarshal([]byte(message), pMsg)

		if pMsg.ID == messages.CHAT_ID {
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

		} else {
			fmt.Printf("Unknown packet received: %+v\n", *pMsg)
			continue
		}
		fmt.Printf("Decoded message: %+v\n", *pMsg)
		sendToLobby <- *pMsg
	}
}

func ActivateReceiver(messageProcessing chan messages.ProcessedMessage, client connection.Connector) {
	defer client.Connection.Close()

	messageTransfer := make(chan string)
	go messageInterpreter(messageTransfer, messageProcessing)

	for client.Scanner.Scan() {
		fmt.Printf("Received message from client %d: %+v\n", client.ConnectorID, client.Scanner.Text())
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
		fmt.Printf("Sending to client %d: %s\n", client.ConnectorID, jsonString)
		client.Connection.Write([]byte(jsonString + "\n"))
	}
}

func ClientListener(lm *lobbyMap.LobbyMap, client connection.Connector) {
	core := new(ClientCore)
	core.client = client
	core.lm = lm
	processedChan := make(chan messages.ProcessedMessage)

	finJSON := make(chan string)

	go ActivateReceiver(processedChan, core.client)
	go ActivateSender(finJSON, core.client)

	refreshList := ReqUpdate(messages.UpdateRooms{}, *core)
	finJSON <- encoders.EncodeRefreshList(messages.REFRESH_ID, refreshList)

	gameChan := make(chan messages.ProcessedMessage)

	for {
		processed := <- processedChan
		fmt.Printf("Received decoded message: %v\n", processed)

		if processed.ID == messages.CHAT_ID {
			gameChan <- processed
	
		} else if processed.ID == messages.HOST_ID {
			hostedRoom := ReqHost(processed.Host, *core)
			gameChan = hostedRoom.GameChan
			go game.CreateGameRoom(hostedRoom, core.lm)

		} else if processed.ID == messages.JOIN_ID {
			joinedRoom := ReqJoin(processed.Join, *core)
			gameChan = joinedRoom.GameChan
			gameChan <- processed

		} else if processed.ID == messages.REFRESH_ID {
			refreshList := ReqUpdate(processed.Update, *core)
			finJSON <- encoders.EncodeRefreshList(messages.REFRESH_ID, refreshList)

		} else {
			fmt.Println("Something went wrong!")
		}
		fmt.Printf("lm=%+v\n", lm.GetShadow())
	}
}
