// Package encoders contains functions to encode packages in JSON.
package encoders

import (
	"encoding/json"
	"server/messages"
)

// EncodeRoom encodes needed info about a gameroom to JSON. Returns a JSON string of the room.
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

// EncodeKick encodes needed info about a kicked client to JSON. Returns a JSON string.
func EncodeKick() string {
	obj := messages.Leave{
		PacketID: messages.KICK_ID,
	}
	objStr, _ := json.Marshal(obj)
	return string(objStr)
}
