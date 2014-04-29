package database

import "server/connection"

type getter struct {
	id int
	sendBack chan connection.Connector
}

type updater struct {
	newCon connection.Connector
	id int
}

type Database struct {
	add, del chan connection.Connector
	get chan getter
	set chan updater
}

func databaseHandler(db *Database, elements map[int]connection.Connector) {
	for {
		select {
		case con := <-db.add:
			elements[con.ConnectorID] = con
		case con := <-db.del:
			delete(elements, con.ConnectorID)
		case getter := <-db.get:
			client := elements[getter.id]
			getter.sendBack <- client
		case newElement := <- db.set:
			elements[newElement.id] = newElement.newCon
		}
	}
}

func New() *Database {
	db := new(Database)
	db.add = make(chan connection.Connector)
	db.del = make(chan connection.Connector)
	db.get = make(chan getter)
	db.set = make(chan updater)
	go databaseHandler(db, make(map[int]connection.Connector))
	return db
}

func (db Database) Add(con connection.Connector) {
	db.add <- con
}

func (db Database) Delete(con connection.Connector) {
	db.del <- con
}

func (db Database) Get(id int) connection.Connector {
	ch := make(chan connection.Connector)
	getter := getter{id, ch}
	db.get <- getter
	return <- ch
}

func (db Database) Update(id int, newElement connection.Connector) {
	db.set <- updater{newElement, id}
}
