package main

import (
	"sync"

	"github.com/project-quiz/quiz-server/channel"
	"github.com/project-quiz/quiz-server/database"
	"github.com/project-quiz/quiz-server/server"
	"github.com/project-quiz/quiz-server/service"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	channelService := channel.NewChannelService()

	gameService := service.NewGameService(channelService)
	go gameService.ListenToJoinGame()
	go gameService.ListenToLeaveGame()

	server := server.New(channelService)
	go server.Start(4500, &wg)

	db, err := database.New()
	go db.TestDBCon(err)

	wg.Wait()
}
