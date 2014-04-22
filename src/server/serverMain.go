// JUST A SIMPLE SERVER START, INITIATING A SIMPLE LISTENER

package main

import (
	"fmt"
	"net"
	"encoding/gob"
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
	Index, Size int64
	clientChannel chan Connector
//	clientComm ClientCommunication
//	readyClients, runningClients  -> map, array or slice?
}

type WaitingHandler struct {
//	localAddr, remAddr net.Addr
	MaxLobbys, InitiatedLobbys int64
//	allClients, idleClients, busyClients  ->     map, array or slice?
}
//------------------------------------------------

//A connector struct (AKA Client. Perhaps rename it)
type Connector struct {
//	connectorID int64
	connection net.Conn
	localAddr, remAddr net.Addr
	//waiter WaitingHandler
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

//Have a channel to created server clients.
//Now make them available to waitingLobbyManager to continue.
//Save in slice, array or map?
func connectionHandler(connectorChannel chan Connector) {
	for {
		client := <- connectorChannel
		fmt.Println(&client.connection)
		go testConnection(client)
	}
}

//Managing which waitinglobby the connectors get to join.
//Handling spawning of new waitinglobbys.
func waitingLobbyManager(lobbyContact chan chan Connector) {
	handler := &WaitingHandler{20, 0}

	connectionChannel := make(chan Connector)
	lobbyContact <- connectionChannel

	go connectionHandler(connectionChannel)
	go initWaitingLobby()
	handler.InitiatedLobbys += 1;
	
}

//A waiting lobby, where two clients will be able to sync up.
//Then get their own GameRoom
func initWaitingLobby() {
	fmt.Println("Creating a Waiting Lobby for the Server");	
//	lobby := &WaitingLobby{1, 10}
	
}

//Request for a place in a waitingpool, so client can ask for a game start and other things
func requestWaitingPlace() {
	
}

//Create a connection. Simulates as a client on the server.
//CREATE A CONNECTOR(Client) PACKAGE FOR THIS?
func initConnector(connection net.Conn, lobbyContact chan Connector) {
	localAddr := connection.LocalAddr()
	remAddr := connection.RemoteAddr()
	connector := &Connector{connection, localAddr, remAddr}

	lobbyContact <- *connector	
}

//Shall only do minimum that's needed for a request to initiate your own very own server
//Purpose is to establish all for connection requests, waiting lobbys, etc.
//Keep now. Redo later..
func main() {
	fmt.Println("start");
	lobbyContact := make(chan chan Connector)
	go waitingLobbyManager(lobbyContact)

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
		go initConnector(conn, clientChannel)
	}	
}
