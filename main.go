package main

import (
	"fmt"

	"guido.arkesteijn/quiz-server/database"
	"guido.arkesteijn/quiz-server/server"
)

var messageID int32 = 1
var stopServer bool = false

func main() {
	srv, err := database.Connect("192.168.2.18", "4600")

	if err != nil {
		fmt.Printf("error", err.Error())
	}

	srv.GetQuestions()
	srv.GetQuestion("test")

	go server.StartServer(4500)

	for {
		if stopServer {
			break
		}
	}
}
