package main

import (
	"fmt"
	"net"
	"bytes"
        "container/list"
	
)
func removeFromList(list *list.List, element net.Conn){
	for e := list.Front(); e != nil; e = e.Next(){
		con := e.Value.(net.Conn)
		if(con == element){
		list.Remove(e)
		}	
	}
	
}

func handleChat (conn net.Conn, connlist *list.List){

	for {

		buf := make([]byte,1024)
		_, err := conn.Read(buf)
		if err != nil {
			conn.Close()
			fmt.Println("A connection has been lost :(")
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

func handleconnection (conn net.Conn, connlist *list.List){
	fmt.Println("Client  connected succesfully")
	connlist.PushBack(conn)
	handleChat(conn, connlist)
}
func main() {
	fmt.Println("Start")
	ln,err := net.Listen("tcp", ":8080")
        connlist := list.New()
	if err != nil {
		// handle error
		fmt.Printf("Error, something went wrong");
	}
	defer ln.Close();
	for {
	conn, err := ln.Accept() // this blocks until connection or error
	if err != nil {
		// handle error
		continue
	}
		go handleconnection(conn, connlist)

	}



}
