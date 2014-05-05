package connection

import (
	"net"
)

type Connector struct {
	ConnectorID int
	Connection net.Conn
	//LocalAddr, RemAddr net.Addr
	//waiter WaitingHandler
}

func clientIdGenerator(sender chan int) {
	id := 0
	for {
		id++
		sender <- id
	}
}

func initServerClient(netConn net.Conn, idChan chan int, lobbyContact chan Connector) {
	client := &Connector{<-idChan, netConn}
	lobbyContact <- *client
}

func requestListener(idChannel chan int, sendClientChannel chan Connector) {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("Error, port listener couldn't be established")
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go initiateServerClient(conn, idChannel, sendClientChannel)
	}
}
