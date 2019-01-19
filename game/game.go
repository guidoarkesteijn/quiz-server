package game

import (
	"fmt"

	model "github.com/project-quiz/quiz-go-model"
)

type Game struct {
	PlayerJoined chan model.PlayerJoin
	Players      []model.Player
}

//New Create new game
func New() (g Game) {
	g = Game{}
	return g
}

//ListenToJoin join a created game.
func (g *Game) ListenToJoin() {
	fmt.Println("Waiting for player")
	p := <-g.PlayerJoined
	fmt.Println(p.Nickname)
}
