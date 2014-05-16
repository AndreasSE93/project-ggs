package encoders

import (
	"encoding/json"
	"server/messages"
)

func EncodeRoom(packetID int, room messages.RoomData) string {
	obj := messages.HostRoomPacket{
		PacketID: packetID,
		HostRoom: room.CS,
	}
	objStr, _ := json.Marshal(obj)
	return string(objStr)
}

func EncodeHostedRoom(room messages.RoomData) string {
	return EncodeRoom(messages.HOST_ID, room)
}

func EncodeJoinedRoom(room messages.RoomData) string {
	return EncodeRoom(messages.JOIN_ID, room)
}

func EncodeKick() string {
	obj := messages.Leave{
		PacketID: messages.KICK_ID,
	}
	objStr, _ := json.Marshal(obj)
	return string(objStr)
}
