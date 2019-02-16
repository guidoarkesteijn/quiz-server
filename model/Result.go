package model

import "github.com/project-quiz/quiz-go-model/message"

type Result struct {
	Message      *message.BaseMessage
	PlayerClient *PlayerClient
}
