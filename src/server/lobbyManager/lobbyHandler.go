package lobbyManager

import (
	"server/connection"
	"server/database"
)

type WaitingLobby struct {
	
}

type HostRoom struct {

}

type JoinRoom struct {
	
}

type UpdateRooms struct {

}

type ChatMessage struct {
	Message string `json:"Message"`
}

type ProcessedMessage struct {
	ID int `json:"ID"`
	ChatM ChatMessage
	Host HostRoom
	Join JoinRoom
	Update UpdateRooms
}

func hostRoom() {

}

func messageInterpreter(messageTransfer chan string, sendToLobby chan ProcessedMessage) {
	for {
		pMsg := &ProcessedMessage{}
		message := <- messageTransfer
		json.UnMarshal([]byte(message), &pMsg)

		if pMsg.ID == 0 {
			chatM := &ChatMessage{}
			json.UnMarshal([]byte(message), &chatM)
			pMsg.ChatM = chatM
			pMsg.Host = nil
			pMsg.Join = nil
			pMsg.Update = nil

		} else if pMsg == 1 {
			hostM := &HostMessage{}
			json.UnMarshal([]byte(message), &hostM)
			pMsg.Host = hostM
			pMsg.chatM = nil
			pMsg.Join = nil
			pMsg.Update = nil

		} else if pMsg == 2 {

		} else if pMsg == 3 {

		} else {
			fmt.Println("ID is malfunctioning")
			return;
		}
		sendToLobby <- pMsg
	}
}

func messageReceiver(messageProcessing chan ProcessedMessage, client connection.Connector) {
	messageTransfer := make(chan string)
	go messageInterpreter(messageTransfer, messageProcessing)

	for {
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			conn.Close()
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
