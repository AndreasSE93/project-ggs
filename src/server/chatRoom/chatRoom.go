// Package chatRoom provides functions for a simple chat room.
package chatRoom

import (
	"fmt"
	"server/connection"
	"server/messages"
	"server/database"
	"server/encoders"
)

// Channels for sending requests to receiver.
type ChatRoom struct {
	add chan connection.Connector
	del chan int
	collectMsg chan messages.ProcessedMessage
	sendCh chan send
}

type send struct {
	msg messages.ProcessedMessage
	shadow []interface{}
}

// SendMessage sends a chat message to all clients connected to the chat room, including the sender.
func (chatRoom ChatRoom) SendMessage(msg messages.ProcessedMessage) {
	chatRoom.collectMsg <- msg
}

// Connect adds a client to the chat room.
// Any messages send to the chat room will be relayed to the new client.
func (chatRoom ChatRoom) Connect(client connection.Connector) {
	chatRoom.add <- client
}

// Disconnect removes a client from the chat room.
// Messages send to the chat room will not be sent to this client any more.
func (chatRoom ChatRoom) Disconnect(client connection.Connector) {
	chatRoom.del <- client.ConnectorID
}

// Go-routine for sending chat messages to all clients.
func sender(outgoing chan send) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Failed to send chat message:", err)
		}
	}()

	for msg := range outgoing {
		for _, e := range msg.shadow {
			conn := e.(connection.Connector)
			fmt.Println(conn.UserName)
			conn.Connection.Write([]byte(encoders.EncodeChatMessage(msg.msg.ChatM, msg.msg.Origin) + "\n"))
		}
	}
}

// Initialise opens a new chat room and returns functions to operate on it.
func Initialise() *ChatRoom {
	db := database.New()
	chatRoom := new(ChatRoom)
	chatRoom.add = make(chan connection.Connector)
	chatRoom.del = make(chan int)
	chatRoom.collectMsg = make(chan messages.ProcessedMessage)
	chatRoom.sendCh = make(chan send)
	
	go receiver(*chatRoom, db)
	go sender(chatRoom.sendCh)
	return chatRoom
}

// Go-routine for listening to requests.
// Used to create a bottleneck to prevent data races.
func receiver(chatRoom ChatRoom, db *database.Database) {
	s := send{}
	for {
		select {
		case e := <- chatRoom.add:
			db.Add(interface{}(e.ConnectorID), interface{}(e))
		case key := <- chatRoom.del:
			db.Delete(interface{}(key))
		case msg := <- chatRoom.collectMsg:
			s.shadow = db.GetShadow()
			s.msg = msg
			chatRoom.sendCh <- s
		}
	}
}
