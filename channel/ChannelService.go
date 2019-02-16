package channel

import (
	"github.com/project-quiz/quiz-go-model/message"
)

//ChannelService contains all channels for communication inside the server.
type ChannelService struct {
	JoinGame   chan message.JoinGame
	GameJoined chan message.GameJoined
}

//NewChannelService Creates a new ChannelService
func NewChannelService() *ChannelService {
	channelService := ChannelService{}
	channelService.JoinGame = make(chan message.JoinGame, 5)
	channelService.GameJoined = make(chan message.GameJoined, 5)
	return &channelService
}
