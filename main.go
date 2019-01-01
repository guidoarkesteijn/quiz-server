package main

import (
	"guido.arkesteijn/quiz-server/server"
)

var messageID int32 = 1
var stopServer bool = false

func main() {
	go server.StartServer(4500)

	for {
		if stopServer {
			break
		}
	}
}
