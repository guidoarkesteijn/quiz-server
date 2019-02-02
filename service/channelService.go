package service

import "github.com/project-quiz/quiz-server/message"

//ChannelService contains all channels for communication inside the server.
type ChannelService struct {
	SendMessageHandler chan message.SendMessage
}

//ChannelService Creates a new ChannelService
func NewChannelService() *ChannelService {
	return &ChannelService{}
}

func (cs ChannelService) Subscribe() {
}
