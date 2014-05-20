package database

import (
	"server/connection"
)

func (db Database) GetID(conn connection.Connector) {

}

func (db Database) GetRoom(conn connection.Connector) int {
	return db.Get(conn.ConnectorID).CurrentRoom
}

func (db Database) SetRoom(conn connection.Connector, newRoomID int) connection.Connector {
	client := db.Get(conn.ConnectorID)
	client.CurrentRoom = newRoomID
	db.Update(client.ConnectorID, client)
	return client
}
