package lobbyManager

import (
	"fmt"
	"bytes"
	"server/connection"
	"server/database/lobbyMap"
	"server/encoders"
	"server/messages"
	"encoding/json"
)

type ProcessedMessage struct {
	ID int `json:"PacketID"`
	ChatM messages.ChatMessage
	Host messages.HostNew
	Join messages.JoinExisting
	Update messages.UpdateRooms
}

type ClientCore struct {
	client connection.Connector
	lm *lobbyMap.LobbyMap
}

func messageInterpreter(messageTransfer chan string, sendToLobby chan ProcessedMessage) {
	for {
		pMsg := new(ProcessedMessage)
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

func ActivateReceiver(messageProcessing chan ProcessedMessage, client connection.Connector) {
	messageTransfer := make(chan string)
	go messageInterpreter(messageTransfer, messageProcessing)

	for {
		buf := make([]byte, 1024)
		_, err := client.Connection.Read(buf)
		if err != nil {
			client.Connection.Close()
			fmt.Printf("Client %d disconnected\n", client.ConnectorID)
			return;
		}
		m := bytes.Index(buf, []byte{0})
		message := string(buf[:m])
		fmt.Printf("Reccived message from client %d: %+v\n", client.ConnectorID, message)
		messageTransfer <- message
	}
}

func ActivateSender(serverTerminal chan string, client connection.Connector) {
	for {
		jsonString := <- serverTerminal
		fmt.Printf("Sending to client %d: %s\n", client.ConnectorID, jsonString)
		client.Connection.Write([]byte(jsonString + "\n"))
	}
}

/*
type test struct {
	UserList [1]string
	GameHost [1]string
	Message string
}
*/

func ClientListener(lm *lobbyMap.LobbyMap, client connection.Connector) {
	core := new(ClientCore)
	core.client = client
	core.lm = lm
	processedChan := make(chan ProcessedMessage)

	finJSON := make(chan string)

	go ActivateReceiver(processedChan, core.client)
	go ActivateSender(finJSON, core.client)

	refreshList := ReqUpdate(messages.UpdateRooms{}, *core)
	finJSON <- encoders.EncodeRefreshList(messages.REFRESH_ID, refreshList)

	for {
		processed := <- processedChan
		fmt.Printf("Received decoded message: %v\n", processed)

		if processed.ID == messages.CHAT_ID {
			finJSON <- encoders.EncodeChatMessage(messages.CHAT_ID, processed.ChatM.Message)
			
		} else if processed.ID == messages.HOST_ID {
			hostedRoom := ReqHost(processed.Host, *core)
			//Some function call to an encoder for when a room is hosted
			//NOT FILLED IN. Look in package encoders
			finJSON <- encoders.EncodeHostedRoom(messages.HOST_ID, hostedRoom)

		} else if processed.ID == messages.JOIN_ID {
			joinedRoom := ReqJoin(processed.Join, *core)
			//Some function call to an encoder for when a room is joined
			//NOT FILLED IN. Look in package encoders
			finJSON <- encoders.EncodeJoinedRoom(messages.JOIN_ID, joinedRoom)

		} else if processed.ID == messages.REFRESH_ID {
			refreshList := ReqUpdate(processed.Update, *core)
			//Some function call to an encoder for when a update is requested
			//NOT FILLED IN. Look in package encoders
			finJSON <- encoders.EncodeRefreshList(messages.REFRESH_ID, refreshList)

		} else {
			fmt.Println("Something went wrong!")
		}
		fmt.Printf("lm=%+v", lm.GetShadow())
	}
}
