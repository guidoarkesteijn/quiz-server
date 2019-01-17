package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/project-quiz/quiz-go-model/quizmodel"
	"github.com/project-quiz/quiz-server/database"
	"github.com/project-quiz/quiz-server/game"
	"github.com/project-quiz/quiz-server/server"
)

var messageID int32 = 1
var stopServer bool = false

func main() {

	channels := CreateChannels()

	playerJoin := <-channels.PlayerJoining

	g := game.New()

	quizmodel.PlayerJoin

	fmt.Println(playerJoin)

	w := Players.PlayerJoin{Nickname: "wdowadmawdo"}

	n := proto.MessageName(&w)

	bytes, _ := proto.Marshal(&w)

	return

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

func (channels *channels) sendMessage(typeurl string, bytes []byte) {

	switch typeurl {
	case "Data.PlayerJoin":
		player := Player.PlayerJoin{}
		err := proto.Unmarshal(bytes, &player)
		fmt.Println("error:", err)
		fmt.Println("message:", player)
		channels.PlayerJoining <- player
	case "Data.Question":
		break
	}
}

type channels struct {
	PlayerJoining chan Players.PlayerJoin
}

func CreateChannels() (c channels) {
	channel := channels{}
	channel.PlayerJoining = make(chan Player.PlayerJoin, 0)

	return c
}
