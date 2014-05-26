// Package encoders contains functions to encode packages in JSON.
package encoders

import (
	"encoding/json"
	"server/messages"
)

func EncodeStartable(startable bool) string {
	obj := messages.Startable{
		PacketID: messages.STARTABLE_ID,
		IsStartable: startable,
	}
	objStr, _ := json.Marshal(obj)
	return string(objStr)
}

func EncodeStartGame(started bool, playerID int) string {
	obj := messages.Started{
		PacketID: messages.START_ID,
		Player: playerID,
		Started: started,
	}
	objStr, _ := json.Marshal(obj)
	return string(objStr)
}
