package main

import (
	"fmt"
	"os"

	"guidoarkesteijn/quiz-server/database"
	"guidoarkesteijn/quiz-server/server"
)

var messageID int32 = 1
var stopServer bool = false

func main() {
	for _, element := range os.Environ() {
		fmt.Println(element)
	}

	password, success := os.LookupEnv("DATABASE_PASSWORD")

	if success {
		fmt.Println("Password=", password)
	} else {
		fmt.Print("no password set")
	}

	return

	srv, err := database.Connect("192.168.2.18", "4600")

	if err != nil {
		fmt.Println("error: ", err.Error())
	}

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

	go server.StartServer(4500)

	for {
		if stopServer {
			break
		}
	}
}
