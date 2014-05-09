package main

import (
	"os"
	"fmt"
	"net"
	"time"
	"bytes"
	"encoding/json"
	"container/list"
	"server/database"
	"server/database/lobbyMap"
	"server/connection"
	"server/lobbyManager"
	"server/messages"
	"server/encoders"
)

/*
func trimNulls(array []byte) []byte {
	for i := 0; i < len(array); i++ {
		if array[i] == 0 {
			return array[0:i]
		}
	}
	return array
}
*/

//Use it as a check that all is completed for a client to know it have a connection to the server
//Otherwise, try tell the client to reconnect to the server
func testConnection(conn connection.Connector) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Connection to client %d failed: %v\n", conn.ConnectorID, err)
			conn.Connection.Close()
		}
	}()

	ping := encoders.EncodePing("Connection test") + "\n"
	fmt.Printf("Sending ping to client %d: %s", conn.ConnectorID, ping)
	conn.Connection.Write([]byte(ping))

	in := make([]byte, 1024)
	response := new(messages.Ping)
	conn.Connection.Read(in)
	in = bytes.Trim(in, "\x00")
	json.Unmarshal(in, response)
	if response.Payload != "Connection test" {
		panic(fmt.Errorf("Received wrong ping Payload!=\"Connection test\": %s -> %+v", in, *response))
	}
	fmt.Printf("Connection to client %d successful with ping %+v\n", conn.ConnectorID, time.Now().Sub(response.TimeStamp))
}

//Have a channel to created server clients.
//Now make them available to waitingLobbyManager to continue.
//Save in slice, array or map?
func connectionHandler(connectorChannel chan connection.Connector, db *database.Database, lm *lobbyMap.LobbyMap, waitingLobby chan connection.Connector, conList *list.List) {
	for {
		client := <- connectorChannel
		fmt.Printf("Client %d connected: %+v\n", client.ConnectorID, client)
		testConnection(client)
		db.Add(client)
		go lobbyManager.ClientListener(lm, client)
	}
}

//Managing which waitinglobby the connectors get to join.
//Handling spawning of new waitinglobbys.
func waitingLobbyManager(lobbyContact chan chan connection.Connector, conList *list.List) {
	db := database.New()

	lm := lobbyMap.Init(db)

	connectionChannel := make(chan connection.Connector)
	lobbyContact <- connectionChannel

	wlChannel := make(chan connection.Connector)
	connectionHandler(connectionChannel, db, lm, wlChannel, conList)
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
			fmt.Printf("Reccived faulty connection: %+v\n", err)
			continue
		}
		// a goroutine handles conn so that the loop can accept other connections
		go initConnector(conn, clientChannel, idChannel, conList)
	}	
}
