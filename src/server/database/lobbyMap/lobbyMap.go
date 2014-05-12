package lobbyMap

import (
	"server/connection"
	"server/database"
	"server/messages"
)

type joiner struct {
	RoomID int
	client connection.Connector
	sendBack chan *messages.HostRoom
}

type LobbyMap struct {
	add chan speaker
	del chan int
	join, kick chan joiner
	getShadow chan chan []messages.HostRoom
	getRoom chan getter
	nextID int
	clientDB *database.Database
}

type speaker struct {
	hr messages.HostRoom
	sendBack chan messages.HostRoom
}

type getter struct {
	id int
	sendBack chan messages.HostRoom
}

func createRoom(s speaker, lm *LobbyMap, hostCollection map[int]messages.HostRoom) {
	s.hr.RoomID = lm.nextID
	lm.nextID++

	if roomID := lm.clientDB.GetRoom(s.hr.Clients[0]); roomID > 0 {
		kicker := joiner{}
		kicker.sendBack = make(chan *messages.HostRoom)
		kicker.RoomID = roomID
		kicker.client = s.hr.Clients[0]
		go kickClient(kicker, lm, hostCollection)
		<- kicker.sendBack
	}

	s.hr.Clients[0] = lm.clientDB.SetRoom(s.hr.Clients[0], s.hr.RoomID)
	hostCollection[s.hr.RoomID] = s.hr
	
	s.sendBack <- s.hr
}

func joinRoom(associate joiner, lm *LobbyMap, hostCollection map[int]messages.HostRoom) {
	room, ok := hostCollection[associate.RoomID]
	if room.ClientCount < room.MaxSize && ok {
		if roomID := lm.clientDB.GetRoom(associate.client); roomID > 0 {
			kicker := joiner{}
			kicker.sendBack = make(chan *messages.HostRoom)
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

func deleteRoom(RoomID int, hostCollection map[int]messages.HostRoom) {
	delete(hostCollection, RoomID)
}

func refreshShadow(sendBack chan []messages.HostRoom, hostCollection map[int]messages.HostRoom) {
	giant := make([]messages.HostRoom, len(hostCollection))
	i := 0
	for m := range hostCollection {
		giant[i] = hostCollection[m]
		i++
	}
	sendBack <- giant[0:i]
}

func kickClient(toKick joiner, lm *LobbyMap, hostCollection map[int]messages.HostRoom) {
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

func getRoom(get getter, hostCollection map[int]messages.HostRoom) {
	get.sendBack <- hostCollection[get.id]
}

func mapHandler(lm *LobbyMap, hostCollection map[int]messages.HostRoom) {
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
		case get := <- lm.getRoom:
			getRoom(get, hostCollection)
		}
	}
}

func GetEmptyHostRoom() messages.HostRoom {
	return *new(messages.HostRoom)
}

func Init(clientDB *database.Database) *LobbyMap {
	hostCollection := make(map[int]messages.HostRoom)
	lm := new(LobbyMap)
	lm.add = make(chan speaker)
	lm.del = make(chan int)
	lm.join = make(chan joiner)
	lm.getShadow = make(chan chan []messages.HostRoom)
	lm.kick = make(chan joiner)
	lm.getRoom = make(chan getter)
	lm.clientDB = clientDB
	lm.nextID = 1
	go mapHandler(lm, hostCollection)
	return lm
}

func (lm LobbyMap) Host(hr messages.HostRoom) messages.HostRoom {
	sendBack := make(chan messages.HostRoom)
	speaker := speaker{hr, sendBack}
	lm.add <- speaker
	return <- sendBack
}

func (lm LobbyMap) Join(id int, client connection.Connector) *messages.HostRoom {
	sendBack := make(chan *messages.HostRoom)
	joiner := joiner{id, client, sendBack}
	lm.join <- joiner
	return <- sendBack
}

func (lm LobbyMap) Delete(RoomID int) {
	lm.del <- RoomID
}

func (lm LobbyMap) GetShadow() []messages.HostRoom {
	sendBack := make(chan []messages.HostRoom)
	lm.getShadow <- sendBack
	return <-sendBack
}

func (lm LobbyMap) Kick(conn connection.Connector) *messages.HostRoom {
	kicker := new(joiner)
	kicker.sendBack = make(chan *messages.HostRoom)
	kicker.RoomID = conn.CurrentRoom
	kicker.client = conn
	lm.kick <- *kicker
	return <- kicker.sendBack
}

func (lm LobbyMap) GetRoom(roomID int) messages.HostRoom {
	get := new(getter)
	get.id = roomID
	get.sendBack = make(chan messages.HostRoom)
	lm.getRoom <- *get
	return <- get.sendBack
}
