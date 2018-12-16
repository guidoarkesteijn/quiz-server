package main

import (
	"fmt"

	"github.com/twinj/uuid"

	Data "guido.arkesteijn/quiz-server/data"
	Question "guido.arkesteijn/quiz-server/data/Question"
)

func main() {
	fmt.Println("Starting server")

	baseMessage := Data.BaseMessage{Id: 1345678}
	fmt.Println(baseMessage.Id)

	u := uuid.NewV4()

	var answer1 = Question.Answer{Text: "Answer 1"}
	var answer2 = Question.Answer{Text: "Answer 2"}
	var answer3 = Question.Answer{Text: "Answer 3"}
	var answer4 = Question.Answer{Text: "Answer 4"}

	var answers = []*Question.Answer{&answer1, &answer2, &answer3, &answer4}
	question := Question.Question{Guid: u.String(), Answers: answers}

	fmt.Println(question.Guid)
	for index := 0; index < len(question.Answers); index++ {
		fmt.Println(question.Answers[index].Text)
	}

	fmt.Println("Ended server")
}
