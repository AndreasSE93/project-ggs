package lobbyMap

import (
	"fmt"
	"testing"
	"server/connection"
	"server/database"
	"server/messages"
)

const (
	Sizer = 11
	Joiner = 2
)

func genID(sender chan int) {
	id := 1
	for {
		sender <- id
		id++
	}
}

func createClient(clientIDChan <-chan int, db *database.Database) connection.Connector{
	client := new(connection.Connector)
	client.ConnectorID = <- clientIDChan
	client.UserName = "Tester " + string([]byte{byte('0' + client.ConnectorID)})
	fmt.Println("ADDING TO DB:", client)
	db.Add(interface{}(client.ConnectorID), interface{}(client))
	return *client
}

func hostRooms(clientIDChan chan int, lm *LobbyMap, db *database.Database) {
	for i:=1; i < Sizer; i++ {
		hr := new(messages.ClientSection)
		hr.RoomName = "Test room"
		c := make([]connection.Connector, 2)
		c[0] = createClient(clientIDChan, db)
		hr.Clients = c
		fmt.Printf("HOSTING: %+v\n", *hr)
		r := lm.Host(messages.RoomData{
			CS: *hr,
		})
		fmt.Printf("HOSTED: %+v\nHOSTRET: %+v\n", *hr, r)
	}
}

func joinRooms(clientIDChan chan int, lm *LobbyMap, db *database.Database) {
	for i:=1; i < Sizer; i++ {
		for j := 0; j < Joiner; j++ {
			client := createClient(clientIDChan, db)
			lm.Join(i, client)
		}
	}
}

func TestHost(t *testing.T) {
	db := database.New()
	lm := New(db)

	gid := make(chan int)
	go genID(gid)
	hostRooms(gid, lm, db)
	joinRooms(gid, lm, db)
	shadow := lm.GetShadow()
	for _, room := range shadow {
		if room.CS.RoomID != room.CS.Clients[0].ConnectorID {
			t.Error(room.CS.RoomID, "(roomID) !=", room.CS.Clients[0].ConnectorID, "(ConnectorID)")
		}
		if (room.CS.RoomID - 1) * Joiner + Sizer != room.CS.Clients[1].ConnectorID {
			t.Error(room.CS.RoomID, "(Wrong joiner)")
		}
	}
	t.Error(shadow)
	for _, room := range shadow {
		lm.Kick(room.CS.Clients[0])
	}
//	t.Error(shadow)
}
