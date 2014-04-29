package connection

import "net"

type Connector struct {
	ConnectorID int
	Connection net.Conn
	LocalAddr, RemAddr net.Addr
	//waiter WaitingHandler
}
