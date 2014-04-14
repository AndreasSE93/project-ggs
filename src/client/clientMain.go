// JUST A SIMPLE CLIENT START, SENDING SOMETHING TO SERVER

package main

import (
    "fmt"
    "log"
    "net"
    "encoding/gob"
)

type P struct {
    M, N int64
}

type Noob struct {
	Message string
}

func connectToServer() net.Conn {
	fmt.Println("trying to connect");
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Connection error", err)
	} else {
		fmt.Println("Connected Successfully");
		return conn
	}
	return nil
}

func main() {
	fmt.Println("start client");
	conn := connectToServer()
	if conn != nil {
		encoder := gob.NewEncoder(conn)
		message := &Noob{"Hello I'm client"}
		p:= &P{1,2}
		encoder.Encode(p) 
		encoder.Encode(message)
		dec := gob.NewDecoder(conn)
		m := &Noob{}
		dec.Decode(m)
		fmt.Printf("Server responds %+v to client\n", *m);
		conn.Close()
	}
}
