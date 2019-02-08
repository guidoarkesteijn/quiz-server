package service

import (
	"github.com/project-quiz/quiz-go-model/model"
)

//ChannelService contains all channels for communication inside the server.
type ChannelService struct {
	JoinGame   chan model.JoinGame
	GameJoined chan model.GameJoined
}

//NewChannelService Creates a new ChannelService
func NewChannelService() *ChannelService {
	channelService := ChannelService{}
	channelService.JoinGame = make(chan model.JoinGame, 5)
	channelService.GameJoined = make(chan model.GameJoined, 5)
	return &channelService
}
