package model

import (
	"net"

	"github.com/project-quiz/quiz-server/interpreter"

	"github.com/golang/protobuf/proto"

	"github.com/project-quiz/quiz-go-model/message"
)

//PlayerClient this contains the data for an existing connection.
type PlayerClient struct {
	Guid     string
	NickName string
	Con      net.Conn
}

func (c *PlayerClient) ToProto() *message.Player {
	return &message.Player{Guid: c.Guid, Nickname: c.NickName}
}

func (p *PlayerClient) WriteMessage(message proto.Message) {
	bytes, err := interpreter.WriteBaseMessage(message)

	if err == nil {
		p.Write(bytes)
	}
}

//WriteWithIndex internal use only! To connection with the index and a proto message interface
func (p *PlayerClient) Write(bytes []byte) {
	p.Con.Write(bytes)
}
