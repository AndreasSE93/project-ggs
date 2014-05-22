package chatRoom

import(
	"fmt"
	"server/connection"
	"server/messages"
	"server/database"
	"server/encoders"
)

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

func (chatRoom ChatRoom) SendMessage(msg messages.ProcessedMessage) {
	chatRoom.collectMsg <- msg
}

func (chatRoom ChatRoom) Connect(client connection.Connector) {
	fmt.Println("COnnect")
	chatRoom.add <- client
}

func (chatRoom ChatRoom) Disconnect(client connection.Connector) {
	chatRoom.del <- client.ConnectorID
}

func sender(outgoing chan send) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Player has left chatroom")
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

func receiver(chatRoom ChatRoom, db *database.Database) {
	s := send{}
	for {
		select {
		case e := <- chatRoom.add:
			db.Add(interface{}(e.ConnectorID), interface{}(e))
			fmt.Println(db.GetShadow())
		case key := <- chatRoom.del:
			db.Delete(interface{}(key))
		case msg := <- chatRoom.collectMsg:
			s.shadow = db.GetShadow()
			fmt.Println(s.shadow)
			s.msg = msg
			chatRoom.sendCh <- s
		}
	}
}
