package server

import (
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
)

// Server struct
type Server struct {
	list map[string]net.Conn
	mux  sync.Mutex
}

// New server creation
func New() Server {
	log.Println("Firing up server!")
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
		log.Println("logged |----|  " + conn.RemoteAddr().String())
		s.mux.Unlock()
		return true
	}
	log.Println(conn.RemoteAddr().String() + " already registered")
	return false
}

func (s *Server) unregister(conn net.Conn) {
	s.mux.Lock()
	delete(s.list, conn.RemoteAddr().String())
	log.Println("unlogged |----|  " + conn.RemoteAddr().String())
	log.Println("quit command received. Bye.")
	s.mux.Unlock()
}

func (s *Server) messageReader(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	for {
		//read buff
		blen, err := conn.Read(buffer)
		message := string(buffer[:blen])

		if message == "/quit" {
			s.unregister(conn)
			return
		}

		if blen > 0 && !strings.ContainsAny(message, "EOF") {
			fmt.Println(message)
			s.broadcast(conn, message)
		}

		if err != nil && !strings.ContainsAny(err.Error(), "EOF") {
			log.Println(err)
			return
		}
	}
}

func (s *Server) broadcast(from net.Conn, message string) {
	for address, to := range s.list {
		if to != from {
			log.Println("message to: " + address)
			to.Write([]byte(message))
		}
	}
}

// Run Start up the server. Manages join and leave chat
func (s *Server) Run() {
	// Listen on port TCP 2016
	listener, err := net.Listen("tcp", ":2016")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		//wait for connection
		conn, err := listener.Accept()
		conn.Write([]byte("welcome to the servar"))

		if err != nil {
			log.Println(err)
		} else {
			s.register(conn)
			go s.messageReader(conn)
		}
	}
}
