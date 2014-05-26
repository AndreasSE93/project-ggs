// Package connection contains info about one client on the server.
package connection

import (
	"net"
	"bufio"
)

type Connector struct {
	ConnectorID, CurrentRoom int
	UserName string
	Connection net.Conn
	Scanner *bufio.Scanner
}
