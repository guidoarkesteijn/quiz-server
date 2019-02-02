package game

import (
	"fmt"

	"github.com/project-quiz/quiz-go-model/model"
	"github.com/twinj/uuid"
)

type Game struct {
	Guid         string
	PlayerJoined chan model.PlayerJoin
	Players      []model.Player
}

//New Create new game
func New() (g Game) {
	g = Game{Guid: uuid.NewV4().String()}
	return g
}

func (g *Game) AddPlayer(player model.Player) {

}

//ListenToJoin join a created game.
func (g *Game) ListenToJoin() {
	fmt.Println("Waiting for player")
	p := <-g.PlayerJoined
	fmt.Println(p.Nickname)
}
