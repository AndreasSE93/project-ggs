package lobbyManager

import (
	"server/database"
	"server/database/lobbyMap"
	"server/connection"
)

type LobbyCore struct {
	db *database.Database
	lm *lobbyMap.LobbyMap
}

func newClientReceiver(incomingClientChan chan connection.Connector, dbHandlers LobbyCore) {
	for {
		client := <- incomingClientChan
		dbHandlers.db.Add(client)
		go ClientListener(dbHandlers.lm, client)
	}
}

func lobbyLoader(incomingClientChan chan connection.Connector, dbHandlers LobbyCore) {
	
	go newClientReceiver(incomingClientChan, dbHandlers)
}

func InitiateLobby(mainChan chan chan connection.Connector) {
	incomingClientChan := make(chan connection.Connector)
	mainChan <- incomingClientChan

	dbHandlers := new(LobbyCore)
	dbHandlers.db = database.New()
	dbHandlers.lm = lobbyMap.Init(dbHandlers.db)

	go lobbyLoader(incomingClientChan, *dbHandlers)
}
