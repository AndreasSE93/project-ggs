package encoders

import (
	"server/database/lobbyMap"
)

func EncodeHostedRoom(packageID int, room lobbyMap.HostRoom) string {
	return "H"
}
