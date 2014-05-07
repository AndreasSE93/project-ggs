package lobbyManager

import (
	"server/connection"
	"server/database/lobbyMap"
)

func ReqHost(hostNew HostNew, clientInfo ClientCore) lobbyMap.HostRoom {
	hr := lobbyMap.GetEmptyHostRoom()
	hr.RoomID = -1
	hr.MaxSize = hostNew.MaxSize
	hr.ClientCount = 1
	hr.RoomName = hostNew.RoomName
	c := make([]connection.Connector, hr.MaxSize)
	c[0] = clientInfo.client
	hr.Clients = c

	hostedRoom := clientInfo.lm.Host(hr)
	return hostedRoom
	
}

func ReqJoin(join JoinExisting, clientInfo ClientCore) *lobbyMap.HostRoom {
	return clientInfo.lm.Join(join.RoomID, clientInfo.client)
}

func ReqUpdate(refresh UpdateRooms, clientInfo ClientCore) []lobbyMap.HostRoom {
	return clientInfo.lm.GetShadow()
}

