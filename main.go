package main

import (
	"github.com/project-quiz/quiz-server/database"
	"github.com/project-quiz/quiz-server/server"
)

var stop bool

func main() {

	db, err := database.New()
	go db.TestDBCon(err)

	server := server.New()
	go server.Start(4500)

	for {
		if stop {
			break
		}
	}
}
