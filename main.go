package main

import (
	"flag"
	"github.com/crmejia/go-chat-server/client"
	"github.com/crmejia/go-chat-server/server"
)

func main() {
	serverMode := flag.Bool("server", false, "Server mode")

	flag.Parse()

	if *serverMode {
		s := server.New()
		s.Run()
	} else {
		client.Run()
	}
}
