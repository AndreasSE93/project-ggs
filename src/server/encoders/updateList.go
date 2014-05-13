package encoders

import (
	"fmt"
	"encoding/json"
	"server/messages"
)

func EncodeRefreshList(packageID int, updateList []messages.RoomData) string {

	a := make([]messages.ClientSection, len(updateList))
	for room := range updateList {
		a[room] = updateList[room].CS
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