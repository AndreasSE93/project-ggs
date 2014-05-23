package lobbyMap

import (
	"fmt"
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

func createRoom(s speaker, lm *LobbyMap, hostCollection map[int]messages.RoomData) {
	s.rd.CS.RoomID = lm.nextID
	lm.nextID++

	if roomID := lm.clientDB.GetRoom(s.rd.CS.Clients[0]); roomID > 0 {
		fmt.Println("KIIIIIIIICKED")
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

func deleteRoom(RoomID int, hostCollection map[int]messages.RoomData) {
	delete(hostCollection, RoomID)
}

func refreshShadow(sendBack chan []messages.RoomData, hostCollection map[int]messages.RoomData) {
	giant := make([]messages.RoomData, len(hostCollection))
	i := 0
	for m := range hostCollection {
		giant[i] = hostCollection[m]
		i++
	}
	sendBack <- giant[0:i]
}

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
			fmt.Println("F")
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
			fmt.Println("G")
			
		}
		fmt.Println("H")
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

func getRoom(get getter, hostCollection map[int]messages.RoomData) {
	//fmt.Printf("Getting room %d: %+v\n", get.id, hostCollection[get.id])
	get.sendBack <- hostCollection[get.id]
}

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

func GetEmptyRoomData() messages.RoomData {
	return *new(messages.RoomData)
}

func Init(clientDB *database.Database) *LobbyMap {
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

func (lm LobbyMap) Host(rd messages.RoomData) messages.RoomData {
	sendBack := make(chan messages.RoomData)
	speaker := speaker{rd, sendBack}
	lm.add <- speaker
	return <- sendBack
}

func (lm LobbyMap) Join(id int, client connection.Connector) *messages.RoomData {
	sendBack := make(chan *messages.RoomData)
	joiner := joiner{id, client, sendBack}
	lm.join <- joiner
	return <- sendBack
}

func (lm LobbyMap) Delete(RoomID int) {
	lm.del <- RoomID
}

func (lm LobbyMap) GetShadow() []messages.RoomData {
	sendBack := make(chan []messages.RoomData)
	lm.getShadow <- sendBack
	return <-sendBack
}

func (lm LobbyMap) Kick(conn connection.Connector) *messages.RoomData {

	kicker := new(joiner)
	kicker.sendBack = make(chan *messages.RoomData)
	kicker.RoomID = conn.CurrentRoom
	kicker.client = conn
	lm.kick <- *kicker
	return <- kicker.sendBack
}

func (lm LobbyMap) GetRoom(roomID int) messages.RoomData {
	get := new(getter)
	get.id = roomID
	get.sendBack = make(chan messages.RoomData)
	lm.getRoom <- *get
	return <- get.sendBack
}
