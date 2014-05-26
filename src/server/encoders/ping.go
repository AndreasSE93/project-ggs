// Package encoders contains functions to encode packages in JSON.
package encoders

import (
	"time"
	"encoding/json"
	"server/messages"
)

func EncodePing(payload interface{}) string {
	obj, _ := json.Marshal(messages.Ping{
		PacketID: messages.PING_ID,
		TimeStamp: time.Now(),
		Payload: payload,
	})
	return string(obj)
}
