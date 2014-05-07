package lobbyMap

import (
	//"fmt"
	"testing"
	"server/connection"
)

const (
	Sizer = 10000
	Joiner = 2
)

func TestInit(t *testing.T) {
	initiate := Init()
	t.Log(initiate)
}

func genID(sender chan int) {
	id := 0
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
	for i:=0; i < Sizer; i++ {
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
	for i:=0; i < Sizer; i++ {
		for j := 0; j < Joiner; j++ {
			client := createClient(clientIDChan)
			lm.Join(i, client)
		}
	}
}

func TestHost(t *testing.T) {
	lm := Init()

	gid := make(chan int)
	go genID(gid)
	hostRooms(gid, lm)
	joinRooms(gid, lm)
	shadow := lm.GetShadow()
	for c := range shadow {
		if shadow[c].RoomID != shadow[c].Clients[0].ConnectorID {
			t.Error(shadow[c].RoomID, "Failed")
		}
		if shadow[c].RoomID*Joiner+Sizer != shadow[c].Clients[1].ConnectorID {
			t.Error(shadow[c].RoomID, "Wrong Joiner")
		}
	}
}
