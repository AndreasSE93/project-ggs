package encoders

import (
	"encoding/json"
	"server/messages"
)

func EncodeHostedRoom(packageID int, room messages.RoomData) string {
	obj := messages.HostRoomPacket{
		PacketID: packageID,
		HostRoom: room.CS,
	}
	objStr, _ := json.Marshal(obj)
	return string(objStr)
}
