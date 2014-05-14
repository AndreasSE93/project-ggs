package game

import(
//	"server/messages"
	"server/games/achtung"
)

func InitAchtung(gameRoom *GameRoom) {
	game := new(achtung.Achtung)
	achtung.GenPlayer(game, gameRoom.roomData.CS.Clients)

	for {
//		processed := <- gameRoom.roomChan
	}

}
