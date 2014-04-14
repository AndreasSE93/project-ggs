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

//Structs for waiting stages??
type WaitingMonitor struct {
	Name string
}

type WaitingLobby struct {
	Index int64
}

type WaitingHandler struct {
//	localAddr, remAddr net.Addr
	K int64	
}
//------------------------------------------------

//A connector struct (AKA Client. Perhaps rename it)
type Connector struct {
//	connectorID int64
	connection net.Conn
	localAddr, remAddr net.Addr
	waiter WaitingHandler
}

//Use it as a check that all is completed for a client to know it have a connection to the server
//Otherwise, try tell the client to reconnect to the server
func handleConnection(conn *Connector) {
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

//A waiting lobby, where client two clients will be able to sync up.
//Then get their own GameRoom
func initWaitingLobby() {
	fmt.Println("Creating a Waiting Lobby for the Server");
}

//Request for a place in a waitingpool, so client can ask for a game start and other things
func requestWaitingPlace() {
	
}

//Create a connection. Simulates as a client on the server.
//CREATE A CONNECTOR(Client) PACKAGE FOR THIS?
func initConnector(connection net.Conn, localAddr net.Addr, remAddr net.Addr) *Connector {
	w := &WaitingHandler{10}
	connector := &Connector{connection, localAddr, remAddr, *w}
	
	go handleConnection(connector)
	return connector
}

//Shall only do minimum that's needed for a request to initiate your own very own server
//Purpose is to establish all for connection requests, waiting lobbys, etc.
//Keep now. Redo later..
func main() {
	fmt.Println("start");
	ln, err := net.Listen("tcp", ":8080")
	go initWaitingLobby()
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
		Addr := conn.LocalAddr()
		Remo := conn.RemoteAddr()
		// a goroutine handles conn so that the loop can accept other connections
		go initConnector(conn, Addr, Remo)
	}
}
