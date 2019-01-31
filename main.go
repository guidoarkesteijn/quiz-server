package main

import (
	"sync"

	"github.com/project-quiz/quiz-server/database"
	"github.com/project-quiz/quiz-server/server"
)

var stop bool

func main() {

	db, err := database.New()
	go db.TestDBCon(err)

	server := server.New()

	var wg sync.WaitGroup
	wg.Add(1)
	go server.Start(4500, &wg)

	wg.Wait()
}
