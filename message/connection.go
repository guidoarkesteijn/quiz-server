package message

import (
	"net"

	"github.com/project-quiz/quiz-go-model/model"
)

//Connection this contains the data for an existing connection.
type Connection struct {
	player model.Player
	Index  int
	active bool
	con    net.Conn
}
