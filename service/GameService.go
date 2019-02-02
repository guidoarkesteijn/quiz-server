package service

import (
	"fmt"

	"github.com/project-quiz/quiz-server/game"
	"github.com/project-quiz/quiz-server/message"

	"github.com/project-quiz/quiz-go-model/model"
)

const maxPlayersPerGame = 2

//GameService keeps track of all the running games.
type GameService struct {
	Channels *ChannelService
	games    map[string]game.Game
}

//NewGameService Creates new GameSerice.
func NewGameService(channelService *ChannelService) *GameService {
	return &GameService{Channels: channelService}
}

//Add 's a new game to the game services.
func (gs *GameService) AddPlayer(player model.Player) {
	game := gs.findAvaiableGame()

	game.AddPlayer(player)
	gs.Channels.SendMessageHandler <- message.SendMessage{Indexes: []int{0}, Message: &model.GameJoined{}}
}

//findAvaiableGame should never return nil because a new game is created when all games are full.
func (gs *GameService) findAvaiableGame() game.Game {
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

func (gs *GameService) Get(guid string) game.Game {
	return gs.games[guid]
}
