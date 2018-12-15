package main

import (
	"fmt"

	"github.com/twinj/uuid"

	Data "guido.arkesteijn/quiz-server/data"
)

func main() {
	baseMessage := Data.BaseMessage{}
	fmt.Println(baseMessage.Id)

	u := uuid.NewV4()
	question := Data.Question{Guid: u.String()}

	fmt.Println(question.Guid)
}
