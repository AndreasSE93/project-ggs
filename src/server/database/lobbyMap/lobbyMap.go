package lobbyMap

import (
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
	nextID, groupSize int
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

func createRoom(s speaker, lm *LobbyMap, hostCollection map[int]map[int]HostRoom) {
	s.hr.RoomID = lm.nextID
	lm.nextID++

	groupID := s.hr.RoomID/lm.groupSize
	roomIndex := s.hr.RoomID % lm.groupSize
	if hostCollection[groupID] == nil {
		hostCollection[groupID] = make(map[int]HostRoom)
	}
	hostCollection[groupID][roomIndex] = s.hr
	s.sendBack <- s.hr
}

func joinRoom(associate joiner, lm *LobbyMap, hostCollection map[int]map[int]HostRoom) {
	groupID := associate.RoomID/lm.groupSize
	roomIndex := associate.RoomID % lm.groupSize

	room := hostCollection[groupID][roomIndex]
	if room.ClientCount < room.MaxSize {
		room.ClientCount++
		room.Clients[room.ClientCount] = associate.client
		associate.sendBack <- &room
	} else {
		associate.sendBack <- nil
	}
}

func deleteRoom(RoomID int, lm *LobbyMap, hostCollection map[int]map[int]HostRoom) {
	groupID := RoomID/lm.groupSize
	roomIndex := RoomID % lm.groupSize

	delete(hostCollection[groupID], roomIndex)
}

func refreshShadow(sendBack chan []HostRoom, lm *LobbyMap, hostCollection map[int]map[int]HostRoom) {
	giant := make([]HostRoom, len(hostCollection)*lm.groupSize)
	i := 0
	for m := range hostCollection {
		for e := range hostCollection[m] {
			_, ok := hostCollection[m][e]
			if ok {
				giant[i] = hostCollection[m][e]
				i++
			}
		}
	}
	sendBack <- giant[0:i]
}

func mapHandler(lm *LobbyMap, hostCollection map[int]map[int]HostRoom) {
	for {
		select {
		case speaker := <-lm.add:
			createRoom(speaker, lm, hostCollection)
		case joiner := <-lm.join:
			joinRoom(joiner, lm, hostCollection)
		case del := <-lm.del:
			deleteRoom(del, lm, hostCollection)			
		case getShadow := <-lm.getShadow:
			refreshShadow(getShadow, lm, hostCollection)
		}
	}
}

func Init(roomSize int) *LobbyMap {
	hostCollection := make(map[int]map[int]HostRoom)
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

func (lm LobbyMap) Join(id int, client connection.Connector) HostRoom {
	sendBack := make(chan *HostRoom)
	joiner := joiner{id, client, sendBack}
	lm.join <- joiner
	return *(<- sendBack)
}

func (lm LobbyMap) Delete(RoomID int) {
	lm.del <- RoomID
}

func (lm LobbyMap) GetShadow() []HostRoom {
	sendBack := make(chan []HostRoom)
	lm.getShadow <- sendBack
	return <-sendBack
}
