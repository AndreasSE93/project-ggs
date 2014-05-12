package encoders

import (
	"encoding/json"
	"server/messages"
)

func EncodeHostedRoom(packageID int, room messages.HostRoom) string {
	obj := messages.HostRoomPacket{
		PacketID: packageID,
		HostRoom: room,
	}
	objStr, _ := json.Marshal(obj)
	return string(objStr)
}
