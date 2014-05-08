package encoders

import (
	"encoding/json"
	"server/database/lobbyMap"
	"server/messages"
)

func EncodeRefreshList(packageID int, updateList []lobbyMap.HostRoom) string {
	obj := messages.RoomList{
		PacketID: messages.REFRESH_ID,
		Rooms: updateList,
		Games: []string{"Dummy"},
	}
	objStr, _ := json.Marshal(obj)
	return string(objStr)
}
