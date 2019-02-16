package interpreter

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/project-quiz/quiz-go-model/message"
)

//ReadBaseMessage reads the bytes to a base message.
func ReadBaseMessage(bytes []byte) (message.BaseMessage, error) {
	message := message.BaseMessage{}
	err := proto.Unmarshal(bytes, &message)
	return message, err
}

//WriteBaseMessage writes a proto message inside an base message's any field and returns the marshaled bytes or nil when error occurs.
func WriteBaseMessage(m proto.Message) ([]byte, error) {
	//get name for any field.
	name := proto.MessageName(m)
	serialized, err := proto.Marshal(m)
	if err != nil {
		return nil, err
	}

	anything := &any.Any{
		TypeUrl: name,
		Value:   serialized,
	}

	baseMessage := message.BaseMessage{Message: anything}
	bytes, seconderr := proto.Marshal(&baseMessage)

	if seconderr != nil {
		return nil, seconderr
	}

	return bytes, nil
}
