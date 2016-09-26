package main

import (
	"github.com/crmejia/go-chat-server/server"
)

func main() {
	s := server.New()
	s.Run()
}
