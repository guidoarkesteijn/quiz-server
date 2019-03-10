package server

import (
	"fmt"
	"net"
	"sync"

	"github.com/golang/protobuf/ptypes"
	"github.com/project-quiz/quiz-go-model/message"
	"github.com/project-quiz/quiz-server/channel"
	"github.com/project-quiz/quiz-server/model"
)

var playerClientList map[string]model.PlayerClient

//Service this service contains all code to start an tcp server on the given port create with server.New() and then call func start with port to start it.
type Service struct {
	Channels          *channel.ChannelService
	connected         bool
	onMessageReceived chan model.Result
}

//New create new server service with onMessageReceived channel
func New(channelService *channel.ChannelService) Service {
	playerClientList = make(map[string]model.PlayerClient)
	return Service{onMessageReceived: make(chan model.Result, 10), Channels: channelService}
}

//Start server on the given port.
func (s *Service) Start(port int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Starting server")

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer l.Close()

	fmt.Println("Waiting for clients!")

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		s.connected = true

		go s.WaitForMessage()
		HandleConnection(s, c)
	}
}

//WaitForMessage test function to see how channels work.
func (s *Service) WaitForMessage() {
	for {
		fmt.Println("Waiting for value")
		result := <-s.onMessageReceived

		dynamic := ptypes.DynamicAny{}
		err := ptypes.UnmarshalAny(result.Message.Message, &dynamic)

		if err != nil {
			fmt.Println("error getting any message name:", err.Error())
		} else {
			switch v := dynamic.Message.(type) {
			case *message.JoinServer:
				fmt.Println("Player joins server. Welcome!")
				player := result.PlayerClient
				player.WriteMessage(&message.ServerJoined{Player: player.ToProto()})
			case *message.JoinGame:
				result.PlayerClient.NickName = v.Player.Nickname
				fmt.Println("player wants to join game:", v.Player.Nickname)
				channelService := *s.Channels
				channelService.JoinGame <- *result.PlayerClient
				value := <-s.Channels.GameJoined
				fmt.Println("Found game:", value)
				fmt.Println("Player count:", len(value.Players))
				result.PlayerClient.WriteMessage(&message.GameJoined{GUID: value.Guid, Players: value.ToProto()})
			case *message.PlayerJoin:
				player := result.PlayerClient
				player.NickName = v.Nickname
				player.WriteMessage(&message.PlayerJoined{Guid: player.Guid, Player: player.ToProto()})
			default:
				fmt.Printf("I don't know about this message type. %T!\n", v)
			}
		}
	}
}
