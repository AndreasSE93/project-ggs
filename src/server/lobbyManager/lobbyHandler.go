package lobbyManager

import (
	"fmt"
	"bytes"
	"server/connection"
	"server/database"
	"encoding/json"
)

type HostRoom struct {
	RoomName string `json:"RoomID"`
}

type JoinRoom struct {
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
	Host HostRoom
	Join JoinRoom
	Update UpdateRooms
}

func messageInterpreter(messageTransfer chan string, sendToLobby chan ProcessedMessage) {
	for {
		pMsg := &ProcessedMessage{}
		message := <- messageTransfer
		json.Unmarshal([]byte(message), &pMsg)

		if pMsg.ID == 100 {
			chatM := &ChatMessage{}
			json.Unmarshal([]byte(message), &chatM)
			pMsg.ChatM = *chatM

		} else if pMsg.ID == 101 {
			hostM := &HostRoom{}
			json.Unmarshal([]byte(message), &hostM)
			pMsg.Host = *hostM

		} else if pMsg.ID == 102 {
			joinM := &JoinRoom{}
			json.Unmarshal([]byte(message), &joinM)
			pMsg.Join = *joinM

		} else if pMsg.ID == 103 {
			updateM := &UpdateRooms{}
			json.Unmarshal([]byte(message), &updateM)
			pMsg.Update = *updateM

		} else {
			fmt.Println("ID is malfunctioning")
			return;
		}
		sendToLobby <- *pMsg
	}
}

func messageReceiver(messageProcessing chan ProcessedMessage, client connection.Connector) {
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

func lobbyClientListener(client connection.Connector) {
	processedChan := make(chan ProcessedMessage)

	go messageReceiver(processedChan, client)

	for {
		processed := <- processedChan
		fmt.Println(processed)
	}
}

func newClientHandler(clientFeeder chan connection.Connector, database *database.Database) {
	for {
		client := <- clientFeeder
		database.Add(client)
		go lobbyClientListener(client)
	}
}

func initWaitingLobby(starterChan chan chan connection.Connector, database *database.Database) {
	
	newClientChannel := make(chan connection.Connector)
	starterChan <- newClientChannel

	go newClientHandler(newClientChannel, database)
}
