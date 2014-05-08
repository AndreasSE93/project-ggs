// JUST A SIMPLE SERVER START, INITIATING A SIMPLE LISTENER

package main

import (
	"os"
	"fmt"
	"net"
	"bytes"
	"encoding/gob"
	"encoding/json"
	"container/list"
	"server/database"
	"server/database/lobbyMap"
	"server/connection"
	"server/lobbyManager"
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
	clientChannel chan connection.Connector
//	clientComm ClientCommunication
//	readyClients, runningClients  -> map, array or slice?
}

type WaitingHandler struct {
//	localAddr, remAddr net.Addr
	databaseOp database.Database
	MaxLobbys, InitiatedLobbys int64
//	allClients, idleClients, busyClients []byte
}
//------------------------------------------------

//Use it as a check that all is completed for a client to know it have a connection to the server
//Otherwise, try tell the client to reconnect to the server
func testConnection(conn connection.Connector) {
	dec := gob.NewDecoder(conn.Connection)
	enc := gob.NewEncoder(conn.Connection)
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
	Message string
}

func clientListener(client connection.Connector, conList *list.List) {
//	connection := client.connection
	//enc := json.NewEncoder(client.connection)
	userList := [1]string{"a"}
	gameHost := [1]string{"b"}
	message := "Test"
	b, err := json.Marshal(&test{
		userList,
		gameHost,
		message})
	fmt.Println(string(b), err)
	//buf := []byte(string(b) + "\n")
	c := string(b) + "\n"
	client.Connection.Write([]byte(c))
	handleChat(client.Connection, conList)
}

func handleChat (conn net.Conn, connlist *list.List) {
	fmt.Println(connlist)
	for {

		buf := make([]byte,1024)
		_, err := conn.Read(buf)
		if err != nil {
			conn.Close()
			fmt.Printf("A connection has been lost :(")
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
func connectionHandler(connectorChannel chan connection.Connector, db *database.Database, lm *lobbyMap.LobbyMap, waitingLobby chan connection.Connector, conList *list.List) {
	for {
		client := <- connectorChannel
		fmt.Printf("Client %d connected: %v\n", client.ConnectorID, client)
		db.Add(client)
		//go clientListener(client, conList)
		go lobbyManager.ClientListener(lm, client)
		go testConnection(client)
	}
}

//Managing which waitinglobby the connectors get to join.
//Handling spawning of new waitinglobbys.
func waitingLobbyManager(lobbyContact chan chan connection.Connector, conList *list.List) {
	db := database.New()
	//handler := &WaitingHandler{*database, 20, 0}

	lm := lobbyMap.Init(db)

	connectionChannel := make(chan connection.Connector)
	lobbyContact <- connectionChannel

	wlChannel := make(chan connection.Connector)
	go connectionHandler(connectionChannel, db, lm, wlChannel, conList)

	go initWaitingLobby(wlChannel)
	//handler.InitiatedLobbys += 1;
	
}

//A waiting lobby, where two clients will be able to sync up.
//Then get their own GameRoom
func initWaitingLobby(clientContact chan connection.Connector) {
	fmt.Println("Creating a Waiting Lobby for the Server");	
//	lobbyList := list.New()
//	lobby := &WaitingLobby{1, 10}
}

//Request for a place in a waitingpool, so client can ask for a game start and other things
func requestWaitingPlace() {
	
}

//Create a connection. Simulates as a client on the server.
//CREATE A CONNECTOR(Client) PACKAGE FOR THIS?
func initConnector(netConn net.Conn, lobbyContact chan connection.Connector, idCh chan int, conList *list.List) {
	conn := new(connection.Connector)
	conn.ConnectorID = <- idCh
	conn.Connection  = netConn
	conn.CurrentRoom = -1
	conList.PushBack(conn.Connection)

	lobbyContact <- *conn
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
	lobbyContact := make(chan chan connection.Connector)
	idChannel := make(chan int)

	conList := list.New()

	go waitingLobbyManager(lobbyContact, conList)
	go idGenerator(idChannel)

	//Here to not depend on the channel created in the initiation of the server.
	clientChannel := <-lobbyContact

	listenAddr := ":8080"
	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		fmt.Println("Unable to host server on " + listenAddr)
		os.Exit(1)
	}
	for {
		conn, err := ln.Accept() // this blocks until connection or error
		if err != nil {
			fmt.Printf("Reccived faulty connection: %v\n", err)
			continue
		}
		// a goroutine handles conn so that the loop can accept other connections
		go initConnector(conn, clientChannel, idChannel, conList)
	}	
}
