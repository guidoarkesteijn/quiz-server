package interpreter

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/project-quiz/quiz-go-model/model"
)

//ReadBaseMessage reads the bytes to a base message.
func ReadBaseMessage(bytes []byte) (model.BaseMessage, error) {
	message := model.BaseMessage{}
	err := proto.Unmarshal(bytes, &message)
	return message, err
}

//WriteBaseMessage writes a proto message inside an base message's any field and returns the marshaled bytes or nil when error occurs.
func WriteBaseMessage(message proto.Message) ([]byte, error) {
	//get name for any field.
	name := proto.MessageName(message)
	serialized, err := proto.Marshal(message)
	if err != nil {
		return nil, err
	}

	anything := &any.Any{
		TypeUrl: name,
		Value:   serialized,
	}

	baseMessage := model.BaseMessage{Message: anything}
	bytes, seconderr := proto.Marshal(&baseMessage)

	if seconderr != nil {
		return nil, seconderr
	}

	return bytes, nil
}
