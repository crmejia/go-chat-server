package server

import (
	"fmt"
	"log"
	"net"
	"sync"
)

type Server struct {
	list map[string]net.Conn
	mux  sync.Mutex
}

func New() Server {
	return Server{list: make(map[string]net.Conn)}
}

func (s *Server) register(conn net.Conn) {
	s.mux.Lock()
	s.list[conn.RemoteAddr().String()] = conn
	s.mux.Unlock()
}

func (s *Server) unregister(conn net.Conn) {
	s.mux.Lock()
	delete(s.list, conn.RemoteAddr().String())
	s.mux.Unlock()
}

func (s *Server) messageReader(conn net.Conn) {
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
		fmt.Println("quit command recieved. Bye.")
		//unregister
		s.unregister(conn)
		return
	}
	fmt.Println(message)
}

func (s *Server) Run() {
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
		//register the client
		s.register(conn)
		fmt.Println("registered local " + conn.LocalAddr().String() + "   |----|  remote " + conn.RemoteAddr().String())

		//run goroutines to deal with multiple connections
		go s.messageReader(conn)
	}
}
