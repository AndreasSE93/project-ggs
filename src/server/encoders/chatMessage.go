package encoders

import (
	"encoding/json"
	"server/messages"
)

func EncodeChatMessage(packageID int, msg messages.ChatMessage) string {
	obj, _ := json.Marshal(msg)
	return string(obj)
}
