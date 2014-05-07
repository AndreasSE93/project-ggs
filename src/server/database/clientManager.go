package database

import (
	"fmt"
	"server/connection"
)

func (db Database) GetID(conn connection.Connector) {

}

func (db Database) GetRoom(conn connection.Connector) int {
	return db.Get(conn.ConnectorID).CurrentRoom
}

func (db Database) SetRoom(conn connection.Connector, newRoomID int) connection.Connector {
	fmt.Println(conn.ConnectorID)
	client := db.Get(conn.ConnectorID)
	client.CurrentRoom = newRoomID
	db.Update(client.ConnectorID, client)
//	fmt.Println(client.ConnectorID)
	return client
}
