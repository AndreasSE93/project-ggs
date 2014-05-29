package database

import (
	"fmt"
	"runtime/debug"
	"server/connection"
)

func (db Database) GetRoom(conn connection.Connector) (r int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Error getting room for client %d: %v\n", conn.ConnectorID, err)
			r = -1
			debug.PrintStack()
			println()
			return
		}
	}()
	r = db.Get(interface{}(conn.ConnectorID)).(connection.Connector).CurrentRoom
	return
}

func (db Database) SetRoom(conn connection.Connector, newRoomID int) (client connection.Connector) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Error setting room for client %d: %v\n", conn.ConnectorID, err)
			client = connection.Connector{}
			return
		}
	}()
	client = db.Get(interface{}(conn.ConnectorID)).(connection.Connector)
	client.CurrentRoom = newRoomID
	db.Add(interface{}(client.ConnectorID), interface{}(client))
	return client
}
