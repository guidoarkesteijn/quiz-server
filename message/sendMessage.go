package message

import "github.com/golang/protobuf/proto"

type SendMessage struct {
	Indexes []int
	Message *proto.Message
}
