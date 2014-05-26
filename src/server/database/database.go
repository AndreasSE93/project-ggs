// Package database provides a database to save arbitrary data in.
package database

type getter struct {
	key interface{}
	sendBack chan interface{}
}

type Database struct {
	add chan Element 
	del chan interface{}
	get chan getter
	getShadow chan chan []interface{}
}

type Element struct {
	key interface{}
	value interface{}
}

func databaseHandler(db *Database, elements map[interface{}]interface{}) {
	for {
		select {
		case e := <-db.add:
			elements[e.key] = e.value
		case e := <-db.del:
			delete(elements, e)
		case getter := <-db.get:
			client := elements[getter.key]
			getter.sendBack <- client
		case getShadow := <- db.getShadow:
			refreshShadow(getShadow, elements)
		}
	}
}

func refreshShadow(sendBack chan []interface{}, dbList map[interface{}]interface{}) {
	giant := make([]interface{}, len(dbList))
	i := 0
	for key := range dbList {
		giant[i] = dbList[key]
		i++
	}
	sendBack <- giant[0:i]
}

// New generates a new empty database and returns functions to operate on it.
func New() *Database {
	db := new(Database)
	db.add = make(chan Element)
	db.del = make(chan interface{})
	db.get = make(chan getter)
	db.getShadow = make(chan chan []interface{})
	go databaseHandler(db, make(map[interface{}]interface{}))
	return db
}

// Add adds an element with given key and value. If it doesn't exist already.
func (db Database) Add(key interface{}, value interface{}) {
	db.add <- Element{key, value}
}

// Delete deletes the element connected to the given key, if it exists.
func (db Database) Delete(key interface{}) {
	db.del <- key
}

// Get returns the element connected to the given key.
func (db Database) Get(key interface{}) interface{} {
	ch := make(chan interface{})
	getter := getter{key, ch}
	db.get <- getter
	return <- ch
}

// GetShadow returns a list of all elements in the database.
func (db Database) GetShadow() []interface{} {
	sendBack := make(chan []interface{})
	db.getShadow <- sendBack
	return <- sendBack
}
