package encoders

import (
	"encoding/json"
	"server/messages"
)

func EncodeHostedRoom(room messages.RoomData) string {
	obj := messages.HostRoomPacket{
		PacketID: messages.HOST_ID,
		HostRoom: room.CS,
	}
	objStr, _ := json.Marshal(obj)
	return string(objStr)
}
