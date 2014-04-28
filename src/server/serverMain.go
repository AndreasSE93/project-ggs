// JUST A SIMPLE SERVER START, INITIATING A SIMPLE LISTENER

package main

import (
	"fmt"
	"net"
	"encoding/gob"
	"encoding/json"
	"container/list"
	"bytes"
)

//Trivial test structs
type P struct {
	M, N int64
}

type Noob struct {
	Message string
}
//------------------------------------------------

type ClientCommunication struct {
	ClientOne, ClientTwo chan string
}

//------------------------------------------------

//Structs for waiting stages??
type WaitingMonitor struct {
	Name string
}

type WaitingLobby struct {
	Index, Size int
	clientChannel chan Connector
//	clientComm ClientCommunication
//	readyClients, runningClients  -> map, array or slice?
}

type WaitingHandler struct {
//	localAddr, remAddr net.Addr
	databaseOp Database
	MaxLobbys, InitiatedLobbys int64
//	allClients, idleClients, busyClients []byte
}
//------------------------------------------------

//A connector struct (AKA Client. Perhaps rename it)
type Connector struct {
	connectorID int
	connection net.Conn
	localAddr, remAddr net.Addr
	//waiter WaitingHandler
}

type Getter struct {
	ID int
	sendBack chan Connector
}

type Update struct {
	newConn Connector
	ID int
}

type Database struct {
	add,del chan Connector
	get chan Getter
	update chan Update
}

//Use it as a check that all is completed for a client to know it have a connection to the server
//Otherwise, try tell the client to reconnect to the server
func testConnection(conn Connector) {
	dec := gob.NewDecoder(conn.connection)
	enc := gob.NewEncoder(conn.connection)
	p := &P{}
	e := &Noob{}
	dec.Decode(p)
	dec.Decode(e)
	fmt.Printf("Received : %+v\n %+v\n", *p, *e);

	enc.Encode(&Noob{"I heard you!"})
}

//A game room for two players. Final stage before game start with two players
func initGameRoom(conn net.Conn) {
	
}

type test struct {
	UserList [1]string
	GameHost [1]string
}

func clientListener(client Connector, conList *list.List) {
//	connection := client.connection
	//enc := json.NewEncoder(client.connection)
	userList := [1]string{"a"}
	gameHost := [1]string{"b"}
	b, err := json.Marshal(&test{
		userList,
		gameHost})
	fmt.Println(string(b), err)
	//buf := []byte(string(b) + "\n")
	c := string(b) + "\n"
	client.connection.Write([]byte(c))
	handleChat(client.connection, conList)
}

func handleChat (conn net.Conn, connlist *list.List) {
	for {

		buf := make([]byte,1024)
		_, err := conn.Read(buf)
		if err != nil {
			conn.Close()
			fmt.Println("A connection has been lost :(")
			return;
		}
		n := bytes.Index(buf, []byte{0})
                message := string(buf[:n])
                fmt.Println("sending message " + message)
                
		for e := connlist.Front(); e != nil; e = e.Next(){
			con2 := e.Value.(net.Conn)
			con2.Write(buf[:])
		}
		
	}

}

//Have a channel to created server clients.
//Now make them available to waitingLobbyManager to continue.
//Save in slice, array or map?
func connectionHandler(connectorChannel chan Connector, addToMap chan Connector, waitingLobby chan Connector, conList *list.List) {
	for {
		client := <- connectorChannel
		fmt.Println(client.connectorID)
		addToMap <- client
		go clientListener(client, conList)
		go testConnection(client)
	}
}

func databaseHandler(operations Database) {
	clientMap := make(map[int]Connector)
	for {
		select {
		case conn := <-operations.add:
			clientMap[conn.connectorID] = conn
		case conn := <-operations.del:
			delete(clientMap, conn.connectorID)
		case get := <-operations.get:
			client:= clientMap[get.ID]
			get.sendBack <- client
		case update := <-operations.update:
			clientMap[update.ID] = update.newConn
		}
	}
}

//Managing which waitinglobby the connectors get to join.
//Handling spawning of new waitinglobbys.
func waitingLobbyManager(lobbyContact chan chan Connector, conList *list.List) {
	add := make(chan Connector)
	del := make(chan Connector)
	get := make(chan Getter)
	update := make(chan Update)
	database := &Database{add, del, get, update}
	handler := &WaitingHandler{*database, 20, 0}

	go databaseHandler(handler.databaseOp)

	connectionChannel := make(chan Connector)
	lobbyContact <- connectionChannel

	wlChannel := make(chan Connector)
	go connectionHandler(connectionChannel, database.add, wlChannel, conList)

	go initWaitingLobby(wlChannel)
	handler.InitiatedLobbys += 1;
	
}

//A waiting lobby, where two clients will be able to sync up.
//Then get their own GameRoom
func initWaitingLobby(clientContact chan Connector) {
	fmt.Println("Creating a Waiting Lobby for the Server");	
//	lobbyList := list.New()
//	lobby := &WaitingLobby{1, 10}
}

//Request for a place in a waitingpool, so client can ask for a game start and other things
func requestWaitingPlace() {
	
}

//Create a connection. Simulates as a client on the server.
//CREATE A CONNECTOR(Client) PACKAGE FOR THIS?
func initConnector(connection net.Conn, lobbyContact chan Connector, idCh chan int, conList *list.List) {
	localAddr := connection.LocalAddr()
	remAddr := connection.RemoteAddr()
	connector := &Connector{<-idCh, connection, localAddr, remAddr}
	conList.PushBack(connector)

	lobbyContact <- *connector
}

func idGenerator(idCh chan int) {
	id := 0
	for {
		id++
		idCh <- id
	}
}

//Shall only do minimum that's needed for a request to initiate your own very own server
//Purpose is to establish all for connection requests, waiting lobbys, etc.
//Keep now. Redo later..
func main() {
	fmt.Println("start");
	lobbyContact := make(chan chan Connector)
	idChannel := make(chan int)

	conList := list.New()

	go waitingLobbyManager(lobbyContact, conList)
	go idGenerator(idChannel)

	//Here to not depend on the channel created in the initiation of the server.
	clientChannel := <-lobbyContact
	
	fmt.Println(lobbyContact)
	fmt.Println(clientChannel)

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
		fmt.Printf("Error, something went wrong");
	}
	for {
		conn, err := ln.Accept() // this blocks until connection or error
		if err != nil {
			// handle error
			continue
		}
		// a goroutine handles conn so that the loop can accept other connections
		go initConnector(conn, clientChannel, idChannel, conList)
	}	
}
