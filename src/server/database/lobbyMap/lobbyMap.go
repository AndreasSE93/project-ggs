package lobbyMap

import (
//	"fmt"
	"server/connection"
)

type joiner struct {
	RoomID int
	client connection.Connector
	sendBack chan *HostRoom
}

type LobbyMap struct {
	add chan speaker
	del chan int
	join chan joiner
	getShadow chan chan []HostRoom
	nextID int
}

type HostRoom struct {
	RoomID, MaxSize, ClientCount int
	RoomName string
	Clients []connection.Connector
}

type speaker struct {
	hr HostRoom
	sendBack chan HostRoom
}

func createRoom(s speaker, lm *LobbyMap, hostCollection map[int]HostRoom) {
	s.hr.RoomID = lm.nextID
	lm.nextID++

	hostCollection[s.hr.RoomID] = s.hr
	s.sendBack <- s.hr
}

func joinRoom(associate joiner, lm *LobbyMap, hostCollection map[int]HostRoom) {
	room, ok := hostCollection[associate.RoomID]
	if room.ClientCount < room.MaxSize && ok {
		room.ClientCount++
		room.Clients[room.ClientCount - 1] = associate.client
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
		}
	}
}

func Init() *LobbyMap {
	hostCollection := make(map[int]HostRoom)
	lm := new(LobbyMap)
	lm.add = make(chan speaker)
	lm.del = make(chan int)
	lm.join = make(chan joiner)
	lm.getShadow = make(chan chan []HostRoom)
	lm.nextID = 0
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

func (lm LobbyMap) GetShadow() []HostRoom {
	sendBack := make(chan []HostRoom)
	lm.getShadow <- sendBack
	return <-sendBack
}
