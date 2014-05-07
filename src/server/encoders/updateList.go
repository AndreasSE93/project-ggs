package encoders

import (
	"server/database/lobbyMap"
)

func EncodeRefreshList(packageID int, updateList []lobbyMap.HostRoom) string {
	return "R"
}
