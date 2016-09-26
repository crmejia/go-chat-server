package main

import (
	"fmt"
	"net"
)

func main() {
	addr := net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 2016}
	// Listen on port UDP 2016
	listener, err := net.ListenUDP("udp", &addr)
	if err != nil {
		return
	}
	defer listener.Close()

	for {
		//read buff
		buffer := make([]byte, 1024)
		blen, _, err := listener.ReadFromUDP(buffer)
		if err != nil {
			return
		}
		message := string(buffer[:blen])
		if message == "/quit" {
			fmt.Println("quit message recieved. Bye.")
			return
		}
		fmt.Println(message)
	}

}
