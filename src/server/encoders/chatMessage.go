package encoders

import (
	"encoding/json"
	"server/connection"
	"server/messages"
)

func EncodeChatMessage(packageID int, msg messages.ChatMessage, client connection.Connector) string {
	obj, _ := json.Marshal(messages.ChatMessage{
		PacketID: messages.CHAT_ID,
		Message:  msg.Message,
		User:     client.UserName,
	})
	return string(obj)
}
