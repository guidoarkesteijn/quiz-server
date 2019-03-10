package game

import (
	"fmt"

	"github.com/project-quiz/quiz-go-model/message"
	"github.com/project-quiz/quiz-server/model"
	"github.com/twinj/uuid"
)

type Game struct {
	Guid    string
	Players []*model.PlayerClient
}

//New Create new game
func New() (g Game) {
	g = Game{Guid: uuid.NewV4().String()}
	return g
}

func (g *Game) AddPlayer(player *model.PlayerClient) {
	for _, element := range g.Players {
		element.WriteMessage(&message.PlayerJoined{Guid: player.Guid, Player: player.ToProto()})
	}

	g.Players = append(g.Players, player)
}

func (g *Game) RemovePlayer(player *model.PlayerClient) {
	indexToBeDeleted := -1

	for index, element := range g.Players {
		if element.Guid == player.Guid {
			indexToBeDeleted = index
		}
	}

	fmt.Println("Found index to be deleted: ", indexToBeDeleted)

	if indexToBeDeleted > -1 {
		g.Players = RemoveIndex(g.Players, indexToBeDeleted)

		for _, element := range g.Players {
			fmt.Println("Write Player left message.")
			element.WriteMessage(&message.PlayerLeft{Guid: player.Guid, Player: player.ToProto()})
		}
	}
}

func RemoveIndex(s []*model.PlayerClient, index int) []*model.PlayerClient {
	return append(s[:index], s[index+1:]...)
}

func (g *Game) ToProto() (Player []*message.Player) {
	players := make([]*message.Player, len(g.Players))

	for index, element := range g.Players {
		if element != nil {
			players[index] = element.ToProto()
		}
	}

	return players
}
