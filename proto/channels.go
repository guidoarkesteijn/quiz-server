package proto

import (
	"github.com/project-quiz/quiz-go-model/model"
)

type Channels struct {
	PlayerJoined chan model.PlayerJoin
}
