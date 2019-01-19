package main

import (
	"fmt"

	"github.com/project-quiz/quiz-server/database"
	"github.com/project-quiz/quiz-server/server"
)

var stop bool

func main() {
	db, err := database.New()

	if err != nil {
		fmt.Println("error: ", err.Error())
	} else {
		questions, questionErr := db.GetQuestions()

		if questionErr != nil {
			fmt.Println("error: " + questionErr.Error())
		}

		for _, element := range questions {
			fmt.Println("Question:", element.Question)
			fmt.Println("Answers:")
			for _, answer := range element.Answers {
				fmt.Println(answer.Text)
			}
		}
	}

	server := server.New()
	go server.Start(4500)

	for {
		if stop {
			break
		}
	}
}
