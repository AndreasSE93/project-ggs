package encoders

import (
	"encoding/json"
	"server/messages"
)

func EncodeJoinedRoom(room *messages.RoomData) string {
	obj := messages.HostRoomPacket{
		PacketID: messages.JOIN_ID,
		HostRoom: room.CS,
	}
	objStr, _ := json.Marshal(obj)
	return string(objStr)
}
