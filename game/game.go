package game

import (
	"fmt"

	Data "github.com/project-quiz/quiz-go-model/Player"
)

type Game struct {
	PlayerJoined chan Data.PlayerJoin
	Players      []Data.Player
}

//New Create new game
func New() (g Game) {
	g = Game{}
	return g
}

//ListenToJoin join a created game.
func (g *Game) listenToJoin() {
	fmt.Println("Waiting for player")
	p := <-g.PlayerJoined
	fmt.Println(p.Nickname)
}
