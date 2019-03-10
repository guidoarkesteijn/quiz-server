package service

import (
	"fmt"

	"github.com/project-quiz/quiz-server/model"

	"github.com/project-quiz/quiz-server/channel"

	"github.com/project-quiz/quiz-server/game"
)

const maxPlayersPerGame = 100

//GameService keeps track of all the running games.
type GameService struct {
	Channels *channel.ChannelService
	games    map[string]*game.Game
}

//NewGameService Creates new GameSerice.
func NewGameService(channelService *channel.ChannelService) *GameService {
	gameService := GameService{Channels: channelService}
	gameService.games = make(map[string]*game.Game)
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

func (gs *GameService) ListenToLeaveGame() {
	for {
		value := <-gs.Channels.LeaveGame
		gs.OnLeaveGame(&value)
	}
}

func (gs *GameService) OnLeaveGame(player *model.PlayerClient) {
	fmt.Println("OnLeaveGame", player.Guid)

	for element := range gs.games {
		game := gs.games[element]

		game.RemovePlayer(player)
	}
}

func (gs *GameService) OnJoinGame(player *model.PlayerClient) {
	game := gs.FindAvailableGame()
	game.AddPlayer(player)
	gs.Channels.GameJoined <- *game
}

//FindAvailableGame should never return nil because a new game is created when all games are full.
func (gs *GameService) FindAvailableGame() *game.Game {
	for element := range gs.games {
		game := gs.games[element]

		fmt.Println(len(game.Players))
		if len(game.Players) < maxPlayersPerGame {
			return game
		}

		fmt.Println("no empty game found create new game")
	}

	fmt.Println("Creating new game")

	game := game.New()

	gs.games[game.Guid] = &game

	return &game
}

//Get get the specific game back with the matching guid given.
func (gs *GameService) Get(guid string) *game.Game {
	return gs.games[guid]
}
