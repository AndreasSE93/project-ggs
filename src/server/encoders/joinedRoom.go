package encoders

import (
	"encoding/json"
	"server/messages"
)

func EncodeJoinedRoom(packageID int, room *messages.RoomData) string {
	obj := messages.HostRoomPacket{
		PacketID: packageID,
		HostRoom: room.CS,
		Player:   room.CS.ClientCount,
	}
	objStr, _ := json.Marshal(obj)
	return string(objStr)
}
