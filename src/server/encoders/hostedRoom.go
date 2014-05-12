package encoders

import (
	"encoding/json"
	"server/database/lobbyMap"
	"server/messages"
)

func EncodeHostedRoom(packageID int, room lobbyMap.HostRoom) string {
	obj := messages.HostRoom{
		PacketID: packageID,
		HostRoom: room,
	}
	objStr, _ := json.Marshal(obj)
	return string(objStr)
}
