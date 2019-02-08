package message

import (
	"net"

	"github.com/project-quiz/quiz-server/interpreter"

	"github.com/golang/protobuf/proto"

	"github.com/project-quiz/quiz-go-model/model"
)

//Connection this contains the data for an existing connection.
type Connection struct {
	Guid   string
	Player *model.Player
	Index  int
	Con    net.Conn
}

func (c *Connection) WriteMessage(message proto.Message) {
	bytes, err := interpreter.WriteBaseMessage(message)

	if err == nil {
		c.Con.Write(bytes)
	}
}
