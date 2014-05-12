package connection

import (
	"net"
	"fmt"
)

type Connector struct {
	ConnectorID, CurrentRoom int
	UserName string
	Connection net.Conn
}

func clientIdGenerator(sender chan int) {
	id := 0
	for {
		id++
		sender <- id
	}
}

func initServerClient(netConn net.Conn, idChan chan int, lobbyContact chan Connector) {
	client := new(Connector)
	client.ConnectorID = <- idChan
	client.Connection = netConn
	client.CurrentRoom = -1
	lobbyContact <- *client
}

func requestListener(idChannel chan int, sendClientChannel chan Connector) {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("Error, port listener couldn't be established");
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go initServerClient(conn, idChannel, sendClientChannel)
	}
}
