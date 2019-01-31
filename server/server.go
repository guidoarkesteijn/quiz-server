package server

import (
	"fmt"
	"net"
	"sync"

	"github.com/golang/protobuf/ptypes"
	"github.com/project-quiz/quiz-go-model/model"
	"github.com/project-quiz/quiz-server/channels"
)

var connectionMap map[int]Connection

type Result struct {
	Message    model.BaseMessage
	Connection Connection
}

//Service this service contains all code to start an tcp server on the given port create with server.New() and then call func start with port to start it.
type Service struct {
	Channels          *channels.Channels
	connected         bool
	onMessageReceived chan Result
}

//New create new server service with onMessageReceived channel
func New() Service {
	return Service{onMessageReceived: make(chan Result, 10), Channels: channels.Instance()}
}

//Start start server on the given port.
func (s *Service) Start(port int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Starting server")
	connectionMap = make(map[int]Connection)

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
			fmt.Println("error getting any message name: ", err.Error())
		} else {
			switch v := dynamic.Message.(type) {
			case *model.PlayerJoin:

				fmt.Println("PlayerJoin Message!", v.Nickname)
				result.Connection.player.Nickname = v.Nickname
				fmt.Println("Connection: ", result.Connection.player.Guid)
				WriteSingle(result.Connection.Index, &model.PlayerJoined{Guid: result.Connection.player.Guid, Player: &result.Connection.player})
			case *model.PlayerJoined:
				fmt.Println("PlayerJoined Message!")
				//s.Channels.PlayerJoinedEvent <- v
			default:
				fmt.Printf("I don't know about this message type. %T!\n", v)
			}
		}
	}
}
