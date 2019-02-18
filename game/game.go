package game

import (
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
	g.Players = append(g.Players, player)
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
