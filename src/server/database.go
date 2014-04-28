package database

import "net"

type Connector struct {
	connectorID int
	connection net.Conn
	localAddr, remAddr net.Addr
	//waiter WaitingHandler
}

type getter struct {
	id int
	sendBack chan Connector
}

type updater struct {
	newCon Connector
	id int
}

type Database struct {
	add, del chan Connector
	get chan getter
	set chan updater
}

func databaseHandler(db *Database, elements map[int]Connector) {
	for {
		select {
		case con := <-db.add:
			elements[con.connectorID] = con
		case con := <-db.del:
			delete(elements, con.connectorID)
		case getter := <-db.get:
			client := elements[getter.id]
			getter.sendBack <- client
		case newElement := <- db.set:
			elements[newElement.id] = newElement.newCon
		}
	}
}

func New() (db *Database) {
	db.add = make(chan Connector)
	db.del = make(chan Connector)
	db.get = make(chan getter)
	db.set = make(chan updater)
	go databaseHandler(db, make(map[int]Connector))
	return
}

func (db Database) Add(con Connector) {
	db.add <- con
}

func (db Database) Delete(con Connector) {
	db.del <- con
}

func (db Database) Get(id int) Connector {
	ch := make(chan Connector)
	getter := getter{id, ch}
	db.get <- getter
	return <- ch
}

func (db Database) Update(id int, newElement Connector) {
	db.set <- updater{newElement, id}
}
