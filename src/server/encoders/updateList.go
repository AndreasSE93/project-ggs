package encoders

import (
	"fmt"
	"encoding/json"
	"server/messages"
)

func EncodeRefreshList(packageID int, updateList []messages.HostRoom) string {

	a := make([]messages.RoomInfo, len(updateList))
	for room := range updateList {
		a[room] = messages.RoomInfo{
			RoomID: updateList[room].RoomID,
			MaxSize: updateList[room].MaxSize,
			ClientCount: updateList[room].ClientCount,
			RoomName: updateList[room].RoomName,
			GameName: updateList[room].GameName,
		}
	}

	obj := messages.RoomList{
		PacketID: messages.REFRESH_ID,
		Rooms: a,
		Games: []string{"Dummy"},
	}
	objStr, err := json.Marshal(obj)
	if err != nil {
		fmt.Println(err,"Error in EncodeRefreshList")
	}
	return string(objStr)
}
