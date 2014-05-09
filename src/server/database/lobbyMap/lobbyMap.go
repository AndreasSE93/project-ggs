package lobbyMap

import (
	"server/connection"
	"server/database"
)

type joiner struct {
	RoomID int
	client connection.Connector
	sendBack chan *HostRoom
}

type LobbyMap struct {
	add chan speaker
	del chan int
	join, kick chan joiner
	getShadow chan chan []HostRoom
	nextID int
	clientDB *database.Database
}

type HostRoom struct {
	RoomID, MaxSize, ClientCount int
	RoomName string
	Clients []connection.Connector
	GameName string
}

type speaker struct {
	hr HostRoom
	sendBack chan HostRoom
}

func createRoom(s speaker, lm *LobbyMap, hostCollection map[int]HostRoom) {
	s.hr.RoomID = lm.nextID
	lm.nextID++

	if roomID := lm.clientDB.GetRoom(s.hr.Clients[0]); roomID > 0 {
		kicker := joiner{}
		kicker.sendBack = make(chan *HostRoom)
		kicker.RoomID = roomID
		kicker.client = s.hr.Clients[0]
		go kickClient(kicker, lm, hostCollection)
		<- kicker.sendBack
	}

	s.hr.Clients[0] = lm.clientDB.SetRoom(s.hr.Clients[0], s.hr.RoomID)
	hostCollection[s.hr.RoomID] = s.hr
	
	s.sendBack <- s.hr
}

func joinRoom(associate joiner, lm *LobbyMap, hostCollection map[int]HostRoom) {
	room, ok := hostCollection[associate.RoomID]
	if room.ClientCount < room.MaxSize && ok {
		if roomID := lm.clientDB.GetRoom(associate.client); roomID > 0 {
			kicker := joiner{}
			kicker.sendBack = make(chan *HostRoom)
			kicker.RoomID = roomID
			kicker.client = associate.client
			go kickClient(kicker, lm, hostCollection)
			<- kicker.sendBack
		}
		room.ClientCount++
		room.Clients[room.ClientCount - 1] = lm.clientDB.SetRoom(associate.client, room.RoomID)
		hostCollection[room.RoomID] = room
		associate.sendBack <- &room
	} else {
		associate.sendBack <- nil
	}
}

func deleteRoom(RoomID int, hostCollection map[int]HostRoom) {
	delete(hostCollection, RoomID)
}

func refreshShadow(sendBack chan []HostRoom, hostCollection map[int]HostRoom) {
	giant := make([]HostRoom, len(hostCollection))
	i := 0
	for m := range hostCollection {
		giant[i] = hostCollection[m]
		i++
	}
	sendBack <- giant[0:i]
}

func kickClient(toKick joiner, lm *LobbyMap, hostCollection map[int]HostRoom) {
	room, ok := hostCollection[lm.clientDB.GetRoom(toKick.client)]
	if ok {
		found := false
		for key := range room.Clients {
			if room.Clients[key].ConnectorID == toKick.client.ConnectorID || found {
				found = true
				if key < len(room.Clients)-1 {
					room.Clients[key] = room.Clients[key+1]
				} else {
					room.Clients[key] = *new(connection.Connector)
				}				
			}
		}
		if found {
			room.ClientCount--
			hostCollection[room.RoomID] = room
			lm.clientDB.SetRoom(toKick.client, -1)
		}
		if room.ClientCount <= 0 {
			deleteRoom(room.RoomID, hostCollection)
			toKick.sendBack <- nil
		} else {
			toKick.sendBack <- &room
		}
	} else {
		toKick.sendBack <- nil
	}
}

func mapHandler(lm *LobbyMap, hostCollection map[int]HostRoom) {
	for {
		select {
		case speaker := <-lm.add:
			createRoom(speaker, lm, hostCollection)
		case joiner := <-lm.join:
			joinRoom(joiner, lm, hostCollection)
		case del := <-lm.del:
			deleteRoom(del, hostCollection)			
		case getShadow := <-lm.getShadow:
			refreshShadow(getShadow, hostCollection)
		case toKick := <-lm.kick:
			kickClient(toKick, lm, hostCollection)
		}
	}
}

func GetEmptyHostRoom() HostRoom {
	return *new(HostRoom)
}

func Init(clientDB *database.Database) *LobbyMap {
	hostCollection := make(map[int]HostRoom)
	lm := new(LobbyMap)
	lm.add = make(chan speaker)
	lm.del = make(chan int)
	lm.join = make(chan joiner)
	lm.getShadow = make(chan chan []HostRoom)
	lm.kick = make(chan joiner)
	lm.clientDB = clientDB
	lm.nextID = 1
	go mapHandler(lm, hostCollection)
	return lm
}

func (lm LobbyMap) Host(hr HostRoom) HostRoom {
	sendBack := make(chan HostRoom)
	speaker := speaker{hr, sendBack}
	lm.add <- speaker
	return <- sendBack
}

func (lm LobbyMap) Join(id int, client connection.Connector) *HostRoom {
	sendBack := make(chan *HostRoom)
	joiner := joiner{id, client, sendBack}
	lm.join <- joiner
	return <- sendBack
}

func (lm LobbyMap) Delete(RoomID int) {
	lm.del <- RoomID
}

func (lm LobbyMap) Kick(conn connection.Connector) *HostRoom {
	kicker := new(joiner)
	kicker.sendBack = make(chan *HostRoom)
	kicker.RoomID = conn.CurrentRoom
	kicker.client = conn
	lm.kick <- *kicker
	return <- kicker.sendBack
}

func (lm LobbyMap) GetShadow() []HostRoom {
	sendBack := make(chan []HostRoom)
	lm.getShadow <- sendBack
	return <-sendBack
}
