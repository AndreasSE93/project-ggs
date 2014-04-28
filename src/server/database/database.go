package database

import "server/connector"

type getter struct {
	id int
	sendBack chan connector.Connector
}

type updater struct {
	newCon connector.Connector
	id int
}

type Database struct {
	add, del chan connector.Connector
	get chan getter
	set chan updater
}

func databaseHandler(db *Database, elements map[int]connector.Connector) {
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
	db.add = make(chan connector.Connector)
	db.del = make(chan connector.Connector)
	db.get = make(chan getter)
	db.set = make(chan updater)
	go databaseHandler(db, make(map[int]connector.Connector))
	return db
}

func (db Database) Add(con connector.Connector) {
	db.add <- con
}

func (db Database) Delete(con connector.Connector) {
	db.del <- con
}

func (db Database) Get(id int) connector.Connector {
	ch := make(chan connector.Connector)
	getter := getter{id, ch}
	db.get <- getter
	return <- ch
}

func (db Database) Update(id int, newElement connector.Connector) {
	db.set <- updater{newElement, id}
}
