package encoders

import (
	"encoding/json"
	"server/connection"
	"server/messages"
)

func EncodeChatMessage(msg messages.ChatMessage, client connection.Connector) string {
	obj, _ := json.Marshal(messages.ChatMessage{
		PacketID: messages.CHAT_ID,
		Message:  msg.Message,
		User:     client.UserName,
	})
	return string(obj)
}

func EncodeMoveMessage(msg messages.MoveMessage) string {
	obj, _ := json.Marshal(msg)
	return string(obj)
}

func EncodeSnakeMessage(packageID int, playerArray []messages.Player, clear bool, hasWon bool, winner string ) string{
	obj, _ := json.Marshal(messages.SnakeMessage{
		PacketID:    packageID,
		PlayerArray: playerArray,
		ClearBoard: clear,
		HasWon: hasWon,
		WinnerName: winner,
	})
	return string(obj)
}
