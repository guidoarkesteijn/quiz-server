package server

import (
	"fmt"
	"net"
	"sync"

	"github.com/golang/protobuf/ptypes"
	"github.com/project-quiz/quiz-go-model/model"
	"github.com/project-quiz/quiz-server/message"
	"github.com/project-quiz/quiz-server/service"
)

var connectionMap map[int]message.Connection

//Service this service contains all code to start an tcp server on the given port create with server.New() and then call func start with port to start it.
type Service struct {
	Channels          *service.ChannelService
	connected         bool
	onMessageReceived chan message.Result
}

//New create new server service with onMessageReceived channel
func New(channelService *service.ChannelService) Service {
	return Service{onMessageReceived: make(chan message.Result, 10), Channels: channelService}
}

//Start server on the given port.
func (s *Service) Start(port int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Starting server")
	connectionMap = make(map[int]message.Connection)

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
			case *model.PlayerJoin:
				player := result.Connection.Player
				player.Nickname = v.Nickname
				WriteSingle(result.Connection.Index, &model.PlayerJoined{Guid: player.Guid, Player: player})
			case *model.JoinGame:
				fmt.Println("player wants to join game:", v.Player.Nickname)
				channelService := *s.Channels
				channelService.JoinGame <- *v

				value := <-s.Channels.GameJoined
				fmt.Println("found game:", value)
				WriteSingle(result.Connection.Index, &model.GameJoined{})
			default:
				fmt.Printf("I don't know about this message type. %T!\n", v)
			}
		}
	}
}
