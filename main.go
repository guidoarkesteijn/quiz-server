package main

import (
	"fmt"
	"sync"

	"github.com/project-quiz/quiz-server/database"
	"github.com/project-quiz/quiz-server/server"
	"github.com/project-quiz/quiz-server/service"
)

func main() {
	channelService := service.NewChannelService()

	gameService := service.NewGameService(channelService)

	fmt.Println(gameService)

	db, err := database.New()
	go db.TestDBCon(err)

	server := server.New(channelService)

	var wg sync.WaitGroup
	wg.Add(1)
	go server.Start(4500, &wg)

	wg.Wait()
}
