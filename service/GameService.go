package service

import (
	"fmt"

	"github.com/project-quiz/quiz-go-model/model"

	"github.com/project-quiz/quiz-server/game"
)

const maxPlayersPerGame = 2

//GameService keeps track of all the running games.
type GameService struct {
	Channels *ChannelService
	games    map[string]game.Game
}

//NewGameService Creates new GameSerice.
func NewGameService(channelService *ChannelService) *GameService {
	gameService := GameService{Channels: channelService}
	gameService.games = make(map[string]game.Game)
	return &gameService
}

func (gs *GameService) ListenToJoinGame() {
	for {
		fmt.Println("Waiting for Join game Event")

		value := <-gs.Channels.JoinGame

		fmt.Println("Got Join Game Event")

		gs.OnJoinGame(&value)
	}
}

func (gs *GameService) OnJoinGame(joinGame *model.JoinGame) {
	fmt.Println("Finding game for player that wants to join")
	game := gs.FindAvaiableGame()
	game.AddPlayer(joinGame.Player)

	fmt.Println("Found and added player to game:", game.Guid)
	gs.Channels.GameJoined <- model.GameJoined{}
}

//findAvaiableGame should never return nil because a new game is created when all games are full.
func (gs *GameService) FindAvaiableGame() game.Game {
	for element := range gs.games {
		game := gs.games[element]

		fmt.Println(len(game.Players))
		if len(game.Players) < maxPlayersPerGame {
			return game
		}

		fmt.Println("no empty game found create new game")
	}

	game := game.New()

	gs.games[game.Guid] = game

	return game
}

//Get get the specific game back with the matching guid given.
func (gs *GameService) Get(guid string) game.Game {
	return gs.games[guid]
}
