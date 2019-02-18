package channel

import (
	"github.com/project-quiz/quiz-server/game"
	"github.com/project-quiz/quiz-server/model"
)

//ChannelService contains all channels for communication inside the server.
type ChannelService struct {
	JoinGame   chan model.PlayerClient
	GameJoined chan game.Game
}

//NewChannelService Creates a new ChannelService
func NewChannelService() *ChannelService {
	channelService := ChannelService{}
	channelService.JoinGame = make(chan model.PlayerClient, 5)
	channelService.GameJoined = make(chan game.Game, 5)
	return &channelService
}
