package lobbyManager

import (
	"fmt"
	"server/connection"
	"server/messages"
)

func ReqHost(hostNew messages.HostNew, clientInfo ClientCore) (hostedRoom messages.RoomData, err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%v", rec)
		}
	}()
	rd := messages.RoomData{}
	rd.CS.RoomID = -1
	rd.CS.ClientCount = 1
	rd.CS.RoomName = hostNew.RoomName
	rd.CS.GameName = hostNew.GameName
	rd.CS.GameType = hostNew.GameName

	switch rd.CS.GameType {
	case "TicTacToe":
		rd.CS.MinSize = 2
		rd.CS.MaxSize = 2
	case "Achtung Die Kurve":
		rd.CS.MinSize = 2
		rd.CS.MaxSize = 8
	default:
		panic("Invalid game type: " + rd.CS.GameType)
	}

	rd.SS.GameChan = make(chan messages.ProcessedMessage)
	c := make([]connection.Connector, rd.CS.MaxSize)
	c[0] = clientInfo.client
	rd.CS.Clients = c

	hostedRoom = clientInfo.lm.Host(rd)
	return
	
}

func ReqJoin(join messages.JoinExisting, clientInfo ClientCore) *messages.RoomData {
	return clientInfo.lm.Join(join.RoomID, clientInfo.client)
}

func ReqUpdate(refresh messages.UpdateRooms, clientInfo ClientCore) []messages.RoomData {
	return clientInfo.lm.GetShadow()
}

