package main

import (
	"fmt"

	"github.com/project-quiz/quiz-server/database"
	"github.com/project-quiz/quiz-server/server"
)

var messageID int32 = 1
var stopServer bool = false

func main() {
	srv, err := database.Connect()

	if err != nil {
		fmt.Println("error: ", err.Error())
	} else {
		questions, questionErr := srv.GetQuestions()

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

	go server.StartServer(4500)

	for {
		if stopServer {
			break
		}
	}
}
