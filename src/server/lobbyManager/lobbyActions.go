package lobbyManager

import (
	"server/connection"
	"server/database/lobbyMap"
	"server/messages"
)

func ReqHost(hostNew messages.HostNew, clientInfo ClientCore) lobbyMap.HostRoom {
	hr := lobbyMap.GetEmptyHostRoom()
	hr.RoomID = -1
	hr.MaxSize = hostNew.MaxSize
	hr.ClientCount = 1
	hr.RoomName = hostNew.RoomName
	hr.GameName = hostNew.GameName
	c := make([]connection.Connector, hr.MaxSize)
	c[0] = clientInfo.client
	hr.Clients = c

	hostedRoom := clientInfo.lm.Host(hr)
	return hostedRoom
	
}

func ReqJoin(join messages.JoinExisting, clientInfo ClientCore) *lobbyMap.HostRoom {
	return clientInfo.lm.Join(join.RoomID, clientInfo.client)
}

func ReqUpdate(refresh messages.UpdateRooms, clientInfo ClientCore) []lobbyMap.HostRoom {
	return clientInfo.lm.GetShadow()
}

