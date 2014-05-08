package encoders

import (
	"encoding/json"
	"server/messages"
)

func EncodeChatMessage(packageID int, msg string) string {
	obj, _ := json.Marshal(messages.ChatMessage{
		PacketID: packageID,
		Message: msg,
	})
	return string(obj)
}
