package server

import (
	"fmt"
	"log"
	"net"
	"sync"
)

// Server struct
type Server struct {
	list map[string]net.Conn
	mux  sync.Mutex
}

// New server creation
func New() Server {
	return Server{list: make(map[string]net.Conn)}
}

// func (s *Server) isRegistered(conn net.Conn) bool {
// 	if s.list[conn.RemoteAddr().String()] == nil {
// 		return false
// 	}
// 	return true
// }

//rather have one register function with the logic to assess wether a client is
//already registered
func (s *Server) register(conn net.Conn) bool {
	if s.list[conn.RemoteAddr().String()] == nil {
		s.mux.Lock()
		s.list[conn.RemoteAddr().String()] = conn
		log.Println("registered remote |----|  " + conn.RemoteAddr().String())
		s.mux.Unlock()
		return true
	}
	log.Println(conn.RemoteAddr().String() + " already registered")
	return false
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
		fmt.Println("quit command received. Bye.")

		s.unregister(conn)
		return
	}
	fmt.Println(message)
	// time.Sleep(100 * time.Millisecond)
}

// Run Start up the server. Manages join and leave chat
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
		if s.register(conn) {
			//run goroutines to deal with multiple connections
			go s.messageReader(conn)
		} else {
			log.Println("Already registed  |----| " + conn.RemoteAddr().String())
		}
	}
}
