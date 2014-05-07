package lobbyManager

import (
	"fmt"
	"server/connection"
	"server/database/lobbyMap"
	"bytes"
	"encoding/json"
	"server/encoders"
)

type HostNew struct {
	RoomName string `json:"RoomName"`
	MaxSize int `json:"MaxSize"`
}

type JoinExisting struct {
	RoomID int `json:"RoomID"`
}

type UpdateRooms struct {
	
}

type ChatMessage struct {
	Message string `json:"Message"`
}

type ProcessedMessage struct {
	ID int `json:"PacketID"`
	ChatM ChatMessage
	Host HostNew
	Join JoinExisting
	Update UpdateRooms
}

type ClientCore struct {
	client connection.Connector
	lm *lobbyMap.LobbyMap
}

func messageInterpreter(messageTransfer chan string, sendToLobby chan ProcessedMessage) {
	for {
		pMsg := new(ProcessedMessage)
		message := <- messageTransfer
		json.Unmarshal([]byte(message), &pMsg)

		if pMsg.ID == 100 {
			chatM := new(ChatMessage)
			json.Unmarshal([]byte(message), chatM)
			pMsg.ChatM = *chatM

		} else if pMsg.ID == 101 {
			hostM := new(HostNew)
			json.Unmarshal([]byte(message), hostM)
			pMsg.Host = *hostM

		} else if pMsg.ID == 102 {
			joinM := new(JoinExisting)
			json.Unmarshal([]byte(message), joinM) 
			pMsg.Join = *joinM

		} else if pMsg.ID == 103 {
			updateM := new(UpdateRooms)
			json.Unmarshal([]byte(message), updateM)
			pMsg.Update = *updateM

		} else {
			fmt.Println("ID is malfunctioning")
			return;
		}
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
			return;
		}
		m := bytes.Index(buf, []byte{0})
		message := string(buf[:m])
		messageTransfer <- message
	}
}

func ActivateSender(serverTerminal chan string, client connection.Connector) {
	for {
		jsonString := <- serverTerminal
		fmt.Println(jsonString)
		//Send encoded jsonPackage to client
	}
}

func ClientListener(lm *lobbyMap.LobbyMap, client connection.Connector) {
	core := new(ClientCore)
	core.client = client
	core.lm = lm
	processedChan := make(chan ProcessedMessage)

	finJSON := make(chan string)

	go ActivateReceiver(processedChan, core.client)
	go ActivateSender(finJSON, core.client)
	
	for {
		processed := <- processedChan

		if processed.ID == 100 {
			fmt.Println("Some chat shit here")

		} else if processed.ID == 101 {
			hostedRoom := ReqHost(processed.Host, *core)
			//Some function call to an encoder for when a room is hosted
			//NOT FILLED IN. Look in package encoders
			finJSON <- encoders.EncodeHostedRoom(101, hostedRoom)

		} else if processed.ID == 102 {
			joinedRoom := ReqJoin(processed.Join, *core)
			//Some function call to an encoder for when a room is joined
			//NOT FILLED IN. Look in package encoders
			finJSON <- encoders.EncodeJoinedRoom(102, joinedRoom)

		} else if processed.ID == 103 {
			refreshList := ReqUpdate(processed.Update, *core)
			//Some function call to an encoder for when a update is requested
			//NOT FILLED IN. Look in package encoders
			finJSON <- encoders.EncodeRefreshList(103, refreshList)

		} else {
			fmt.Println("Something went wrong!")
		}
	}
}
