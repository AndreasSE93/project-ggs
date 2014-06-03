package lobbyMap

import (
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

func createClient(t *testing.T, clientIDChan <-chan int, db *database.Database) connection.Connector {
	client := connection.Connector{}
	client.ConnectorID = <- clientIDChan
	client.UserName = "Tester " + string([]byte{byte('0' + client.ConnectorID)})
	t.Log("ADDING TO DB:", client)
	db.Add(interface{}(client.ConnectorID), interface{}(client))
	return client
}

func hostRooms(t *testing.T, clientIDChan <-chan int, lm *LobbyMap, db *database.Database) {
	for i := 1; i < Sizer; i++ {
		c := make([]connection.Connector, 2)
		c[0] = createClient(t, clientIDChan, db)

		eater := make(chan messages.ProcessedMessage)
		go func() {
			for m := range eater {
				t.Logf("Eater %d ate %d packet\n", i, m.ID)
			}
			t.Logf("Eater %d exited\n", c[0])
		}()

		r := lm.Host(messages.RoomData{
			CS: messages.ClientSection{
				RoomName: "Test room",
				ClientCount: 1,
				MinSize: 2,
				MaxSize: 2,
				Clients: c,
			},
			SS: messages.ServerSection{
				GameChan: eater,
			},	
		})
		t.Log("HOSTED: %+v\n", r)
	}
}

func joinRooms(t *testing.T, clientIDChan chan int, lm *LobbyMap, db *database.Database) {
	for i := 1; i < Sizer; i++ {
		for j := 0; j < Joiner; j++ {
			client := createClient(t, clientIDChan, db)
			lm.Join(i, client)
		}
	}
}

func TestHost(t *testing.T) {
	db := database.New()
	lm := New(db)

	gid := make(chan int)
	go genID(gid)
	hostRooms(t, gid, lm, db)
	joinRooms(t, gid, lm, db)
	shadow := lm.GetShadow()
	
	for _, room := range shadow {
		if room.CS.RoomID == room.CS.Clients[0].ConnectorID {
			t.Log(room.CS.RoomID, "(roomID) OK!")
		} else {
			t.Error(room.CS.RoomID, "(roomID) !=", room.CS.Clients[0].ConnectorID, "(ConnectorID)")
		}
		if (room.CS.RoomID - 1) * Joiner + Sizer == room.CS.Clients[1].ConnectorID {
			t.Log(room.CS.RoomID, "(Joiner) OK!")
		} else {
			t.Error(room.CS.RoomID, "(Wrong joiner)")
		}
	}

	t.Log("Shadow before kicking:", shadow)
	for _, room := range shadow {
		t.Log("Kicking", room.CS.Clients[0])
		lm.Kick(room.CS.Clients[0])
	}
	t.Log("Shadow after kicking: ", shadow)
}
