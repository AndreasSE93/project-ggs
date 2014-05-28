package lobbyMap

import (
	"testing"
	"server/connection"
	"server/database"
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

func createClient(clientIDChan chan int) connection.Connector{
	client := new(connection.Connector)
	client.ConnectorID = <- clientIDChan
	return *client
}

func hostRooms(clientIDChan chan int, lm *LobbyMap) {
	for i:=1; i < Sizer; i++ {
		hr := new(HostRoom)
		hr.MaxSize = 2
		hr.ClientCount = 1
		c := make([]connection.Connector, 2)
		c[0] = createClient(clientIDChan)
		hr.Clients = c
		lm.Host(*hr)
	}
}

func joinRooms(clientIDChan chan int, lm *LobbyMap) {
	for i:=1; i < Sizer; i++ {
		for j := 0; j < Joiner; j++ {
			client := createClient(clientIDChan)
			lm.Join(i, client)
		}
	}
}

func TestHost(t *testing.T) {
	lm := New(database.New())

	gid := make(chan int)
	go genID(gid)
	hostRooms(gid, lm)
	joinRooms(gid, lm)
	shadow := lm.GetShadow()
	for c := range shadow {
		if shadow[c].RoomID != shadow[c].Clients[0].ConnectorID {
			t.Error(shadow[c].RoomID, "Failed")
		}
		if (shadow[c].RoomID-1)*Joiner+Sizer != shadow[c].Clients[1].ConnectorID {
			t.Error(shadow[c].RoomID, "Wrong Joiner")
		}
	}
	t.Error(shadow)
	for k := range shadow {
		lm.Kick(shadow[k].Clients[0])
	}
//	t.Error(shadow)
}
