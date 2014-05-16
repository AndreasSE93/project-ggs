package lobbyManager

import (
	"fmt"
	"server/connection"
	"server/database/lobbyMap"
	"server/messages"
)

func ReqHost(hostNew messages.HostNew, clientInfo ClientCore) messages.RoomData {
	rd := lobbyMap.GetEmptyRoomData()
	rd.CS.RoomID = -1
	rd.CS.MaxSize = hostNew.MaxSize
	rd.CS.ClientCount = 1
	rd.CS.RoomName = hostNew.RoomName
	rd.CS.GameName = hostNew.GameName
	rd.CS.GameType = hostNew.GameName
	rd.SS.GameChan = make(chan messages.ProcessedMessage)
	c := make([]connection.Connector, rd.CS.MaxSize)
	c[0] = clientInfo.client
	rd.CS.Clients = c

	hostedRoom := clientInfo.lm.Host(rd)
	return hostedRoom
	
}

func ReqJoin(join messages.JoinExisting, clientInfo ClientCore) *messages.RoomData {
	fmt.Println("ROOM ID", join.RoomID)
	return clientInfo.lm.Join(join.RoomID, clientInfo.client)
}

func ReqUpdate(refresh messages.UpdateRooms, clientInfo ClientCore) []messages.RoomData {
	return clientInfo.lm.GetShadow()
}

