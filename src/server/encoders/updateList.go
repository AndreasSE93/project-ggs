package encoders

import (
	"fmt"
	"encoding/json"
	"server/messages"
)

func EncodeRefreshList(packageID int, updateList []messages.HostRoom) string {

	a := make([]messages.RoomInfo, len(updateList))
	fmt.Println(updateList)
	for room := range updateList {
//		fmt.Println("ROOOOOOOOM",room)
		fmt.Println("NOW", a[room].RoomID)
		a[room] = messages.RoomInfo{
			RoomID: updateList[room].RoomID,
			MaxSize: updateList[room].MaxSize,
			ClientCount: updateList[room].ClientCount,
			RoomName: updateList[room].RoomName,
			GameName: updateList[room].GameName,
		}
//		fmt.Println(a[room], "YOOOOOOOOOOOOOOOoo")
	}

	obj := messages.RoomList{
		PacketID: messages.REFRESH_ID,
		Rooms: a,
		Games: []string{"Dummy"},
	}
	objStr, err := json.Marshal(obj)
	fmt.Println(err,"REFRESH")
	return string(objStr)
}
