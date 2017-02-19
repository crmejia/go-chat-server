package client

import (
	"fmt"
	"log"
	"net"
	"strings"
)

// Send chat
func send() {
	conn, err := net.Dial("tcp", ":2016")

	if err != nil {
		log.Println(err)
	}

	defer conn.Close()
	var message string

	for {
		fmt.Scan(&message)
		conn.Write([]byte(message))
		if message == "./quit" {
			return
		}
	}
}

//Run a client
func Run() {
	log.Println("Go Chat client")

	conn, err := net.Dial("tcp", ":2016")

	if err != nil {
		log.Fatal(err)
	}
	log.Println("local addr " + conn.LocalAddr().String())

	defer conn.Close()
	var message string

	go func() {
		buffer := make([]byte, 1024)
		for {
			//read buff
			blen, err := conn.Read(buffer)
			msg := string(buffer[:blen])

			if blen > 0 && !strings.ContainsAny(msg, "EOF") {
				fmt.Println(msg)
			}

			if err != nil && !strings.ContainsAny(err.Error(), "EOF") {
				log.Println(err)
				return
			}
		}
	}()

	for {
		fmt.Scan(&message)
		conn.Write([]byte(message))
		if message == "/quit" {
			return
		}
	}
}
