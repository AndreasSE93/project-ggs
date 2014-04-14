// JUST A SIMPLE SERVER START, INITIATING A SIMPLE LISTENER

package main

import (
    "fmt"
    "net"
    "encoding/gob"
)

type P struct {
    M, N int64
}
func handleConnection(conn net.Conn) {
    dec := gob.NewDecoder(conn)
    p := &P{}
    dec.Decode(p)
    fmt.Printf("Received : %+v\n", p);
}

func main() {
    fmt.Println("start");
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
        go handleConnection(conn) // a goroutine handles conn so that the loop can accept other connections
    }
}
