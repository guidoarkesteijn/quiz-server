package channels

import (
	"github.com/project-quiz/quiz-go-model/model"
)

//Channels
type Channels struct {
	PlayerJoinedEvent chan model.PlayerJoined
	PlayerLeftEvent   chan model.PlayerLeft
	PlayerJoinEvent   chan model.PlayerJoin
}

var instance *Channels

//Instance get the current instance of the channels.
func Instance() *Channels {
	if instance == nil {
		instance = create()
	}

	return instance
}

func create() *Channels {
	channel := &Channels{
		PlayerJoinedEvent: make(chan model.PlayerJoined, 5),
		PlayerLeftEvent:   make(chan model.PlayerLeft, 5),
		PlayerJoinEvent:   make(chan model.PlayerJoin, 5)}

	return channel
}
