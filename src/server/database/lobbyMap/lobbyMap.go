// Package lobbyMap implements a database for storing and handling lobby rooms.
package lobbyMap

import (
	"server/connection"
	"server/database"
	"server/messages"
)

type joiner struct {
	RoomID int
	client connection.Connector
	sendBack chan *messages.RoomData
}

type LobbyMap struct {
	add chan speaker
	del chan int
	join, kick chan joiner
	getShadow chan chan []messages.RoomData
	getRoom chan getter
	nextID int
	clientDB *database.Database
}

type speaker struct {
	rd messages.RoomData
	sendBack chan messages.RoomData
}

type getter struct {
	id int
	sendBack chan messages.RoomData
}

// Creates a new room based on s.rd with an unique ID and sends it to s.sendBack.
func createRoom(s speaker, lm *LobbyMap, hostCollection map[int]messages.RoomData) {
	s.rd.CS.RoomID = lm.nextID
	lm.nextID++

	if roomID := lm.clientDB.GetRoom(s.rd.CS.Clients[0]); roomID > 0 {
		kicker := joiner{}
		kicker.sendBack = make(chan *messages.RoomData)
		kicker.RoomID = roomID
		kicker.client = s.rd.CS.Clients[0]
		go kickClient(kicker, lm, hostCollection)
		<- kicker.sendBack
	}

	s.rd.CS.Clients[0] = lm.clientDB.SetRoom(s.rd.CS.Clients[0], s.rd.CS.RoomID)
	hostCollection[s.rd.CS.RoomID] = s.rd
	
	s.sendBack <- s.rd
}

// Adds associate.client to room associate.RoomID and sends back the updated room to associate.sendBack.
func joinRoom(associate joiner, lm *LobbyMap, hostCollection map[int]messages.RoomData) {
	room, ok := hostCollection[associate.RoomID]
	if room.CS.ClientCount < room.CS.MaxSize && ok {
		if roomID := lm.clientDB.GetRoom(associate.client); roomID > 0 {
			kicker := joiner{}
			kicker.sendBack = make(chan *messages.RoomData)
			kicker.RoomID = roomID
			kicker.client = associate.client
			go kickClient(kicker, lm, hostCollection)
			<- kicker.sendBack
		}
		room.CS.ClientCount++
		room.CS.Clients[room.CS.ClientCount - 1] = lm.clientDB.SetRoom(associate.client, room.CS.RoomID)
		hostCollection[room.CS.RoomID] = room
		associate.sendBack <- &room
	} else {
		associate.sendBack <- nil
	}
}

// Deletes the room.
func deleteRoom(RoomID int, hostCollection map[int]messages.RoomData) {
	if gameChan := hostCollection[RoomID].SS.GameChan; gameChan != nil {
		close(gameChan)
	}
	delete(hostCollection, RoomID)
}

// Sends a slice of all rooms in the lobby to sendBack.
func refreshShadow(sendBack chan []messages.RoomData, hostCollection map[int]messages.RoomData) {
	giant := make([]messages.RoomData, len(hostCollection))
	i := 0
	for m := range hostCollection {
		giant[i] = hostCollection[m]
		i++
	}
	sendBack <- giant[0:i]
}

// Removes toKick.client from room toKick.RoomID and sends back the updated room to toKick.sendBack.
func kickClient(toKick joiner, lm *LobbyMap, hostCollection map[int]messages.RoomData) {
	room, ok := hostCollection[lm.clientDB.GetRoom(toKick.client)]
	if ok {
		found := false
		for key := range room.CS.Clients {
			if room.CS.Clients[key].ConnectorID == toKick.client.ConnectorID || found {
				found = true
				if key < len(room.CS.Clients)-1 {
					room.CS.Clients[key] = room.CS.Clients[key+1]
				} else {
					room.CS.Clients[key] = *new(connection.Connector)
				}				
			}
		}
		if found {
			//fmt.Printf("Kicked client %d: %+v\n", toKick.client.ConnectorID, room)
			room.CS.ClientCount--
			hostCollection[room.CS.RoomID] = room
			lm.clientDB.SetRoom(toKick.client, -1)
			room.SS.GameChan <- messages.ProcessedMessage{
				ID: messages.JOIN_ID,
				Origin: toKick.client,
				Join: messages.JoinExisting{
					PacketID: messages.JOIN_ID,
					RoomID: room.CS.RoomID,
				},
			}
		}
		if room.CS.ClientCount <= 0 {
			deleteRoom(room.CS.RoomID, hostCollection)
			toKick.sendBack <- nil
		} else {
			toKick.sendBack <- &room
		}
	} else {
		toKick.sendBack <- nil
	}
}

// Sends room getter.RoomID to getter.sendBack.
func getRoom(get getter, hostCollection map[int]messages.RoomData) {
	//fmt.Printf("Getting room %d: %+v\n", get.id, hostCollection[get.id])
	get.sendBack <- hostCollection[get.id]
}

// Go-routine for handling 
func mapHandler(lm *LobbyMap, hostCollection map[int]messages.RoomData) {
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

// New returns a new, empty, lobbyMap.
func New(clientDB *database.Database) *LobbyMap {
	hostCollection := make(map[int]messages.RoomData)
	lm := new(LobbyMap)
	lm.add = make(chan speaker)
	lm.del = make(chan int)
	lm.join = make(chan joiner)
	lm.getShadow = make(chan chan []messages.RoomData)
	lm.kick = make(chan joiner)
	lm.getRoom = make(chan getter)
	lm.clientDB = clientDB
	lm.nextID = 1
	go mapHandler(lm, hostCollection)
	return lm
}

// Host creates and returns a new lobby room based on rd.
// The room will be assigned a new unique ID automatically.
func (lm LobbyMap) Host(rd messages.RoomData) messages.RoomData {
	sendBack := make(chan messages.RoomData)
	speaker := speaker{rd, sendBack}
	lm.add <- speaker
	return <- sendBack
}

// Join adds client too room roomID and returns the updated room.
func (lm LobbyMap) Join(roomID int, client connection.Connector) *messages.RoomData {
	sendBack := make(chan *messages.RoomData)
	joiner := joiner{roomID, client, sendBack}
	lm.join <- joiner
	return <- sendBack
}

// Deletes removes room roomID
func (lm LobbyMap) Delete(roomID int) {
	lm.del <- roomID
}

// GetShadow returns a slice of rooms currently in the lobby.
func (lm LobbyMap) GetShadow() []messages.RoomData {
	sendBack := make(chan []messages.RoomData)
	lm.getShadow <- sendBack
	return <-sendBack
}

// Kick removes conn from the the room conn is currently in.
func (lm LobbyMap) Kick(conn connection.Connector) *messages.RoomData {
	kicker := new(joiner)
	kicker.sendBack = make(chan *messages.RoomData)
	kicker.RoomID = conn.CurrentRoom
	kicker.client = conn
	lm.kick <- *kicker
	return <- kicker.sendBack
}

// GetRoom returns a copy of room roomID.
func (lm LobbyMap) GetRoom(roomID int) messages.RoomData {
	get := new(getter)
	get.id = roomID
	get.sendBack = make(chan messages.RoomData)
	lm.getRoom <- *get
	return <- get.sendBack
}
