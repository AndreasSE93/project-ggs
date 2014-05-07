package encoders

import (
	"server/database/lobbyMap"
)

func EncodeJoinedRoom(packageID int, room *lobbyMap.HostRoom) string {
	return "J"
}
