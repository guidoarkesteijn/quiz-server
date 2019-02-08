package game

import (
	"github.com/project-quiz/quiz-go-model/model"
	"github.com/twinj/uuid"
)

type Game struct {
	Guid    string
	Players []*model.Player
}

//New Create new game
func New() (g Game) {
	g = Game{Guid: uuid.NewV4().String()}
	return g
}

func (g *Game) AddPlayer(player *model.Player) {
	g.Players = append(g.Players, player)
}
