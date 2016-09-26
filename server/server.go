package server

import (
	"fmt"
	"log"
	"net"
)

func Run() {
	// Listen on port TPC 2016
	listener, err := net.Listen("tcp", ":2016")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		//wait for connection
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		//run goroutines to deal with multiple connections
		go messageReader(conn)
	}
}

func messageReader(conn net.Conn) {
	//read buff
	buffer := make([]byte, 1024)
	blen, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	message := string(buffer[:blen])
	// This quits the connetion at least when using netcat
	//TODO look fo the right way to close
	if message == "/quit" {
		fmt.Println("quit message recieved. Bye.")
		return
	}
	fmt.Println(message)
}
