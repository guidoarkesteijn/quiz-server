package message

import (
	"github.com/project-quiz/quiz-go-model/model"
)

type Result struct {
	Message    model.BaseMessage
	Connection Connection
}
