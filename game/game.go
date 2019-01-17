package game

import (
	main "project-quiz/quiz-server/main"
)

type game struct {
	players string
}

//Create Do stuff
func Create() (g game) {
	g := game{}
	return g
}

//Join join a created game.
func (g *game) Join() {

}

func (g *game) AddPlayer(ch main.channels) {

}

//Start do stuff
func Start() {

}
