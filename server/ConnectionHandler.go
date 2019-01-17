package server

import (
	"bufio"
	"fmt"
	"net"

	//TODO: Find the correct use of packages so I don't need to specify the type before the github link.
	BaseMessage "github.com/project-quiz/quiz-go-model"
	Player "github.com/project-quiz/quiz-go-model/Player"
	Welcome "github.com/project-quiz/quiz-go-model/Welcome"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/twinj/uuid"
)

type Connection struct {
	player Player.Player
	Index  int
	active bool
	con    net.Conn
}

var i = 0
var messageCount int32 = 1

func HandleConnection(c net.Conn) {
	player := Player.Player{Guid: uuid.NewV4().String(), Nickname: "<UNKNOWN>"}
	connection := Connection{player, i, true, c}

	connectionMap[i] = connection
	go Read(i, connection)

	fmt.Println("Added connection number:" + fmt.Sprintf("%d", connection.Index))
	i++
}

func Read(index int, connection Connection) {
	c := connection.con

	fmt.Printf("Serving %s\n", c.RemoteAddr().String())

	for {
		fmt.Println("start reading welcome message.")

		//TODO get an idea for which byte I need to delimit for.
		bytes, err := bufio.NewReader(c).ReadBytes(0)

		if err != nil {
			fmt.Println("Error Reading: " + err.Error())
			break
		} else {
			var m BaseMessage.BaseMessage
			changedBytes := bytes[0 : len(bytes)-1]

			err := proto.Unmarshal(changedBytes, &m)

			//TODO decode the correct message found by looking at the m.Message.TypeUrl.
			fmt.Println(proto.MessageType(m.Message.TypeUrl))

			playerJoin := Player.PlayerJoin{}
			err2 := ptypes.UnmarshalAny(m.Message, &playerJoin)

			if err != nil || err2 != nil {
				fmt.Println(err2.Error())
			} else {
				p := &connection.player
				p.Nickname = playerJoin.Nickname

				//Only write to the client that is connecting.
				WriteSingle(index, &Welcome.Welcome{Player: &connection.player})
			}
		}
	}

	fmt.Println("Close connection and remove from connection list :")
	RemoveConnection(index)
	c.Close()
}

//WriteSingle write to one single
func WriteSingle(index int, message proto.Message) {
	indexes := []int{index}
	Write(indexes, message)
}

//Write write to the given indexes
func Write(indexes []int, message proto.Message) {
	for index := 0; index < len(indexes); index++ {
		WriteWithIndex(indexes[index], messageCount, message)
	}
	messageCount++
}

//WriteWithIndex internal use only! To connection with the index and a proto message interface
func WriteWithIndex(index int, messageIndex int32, message proto.Message) {
	name := proto.MessageName(message)

	serialized, err := proto.Marshal(message)
	if err != nil {
		fmt.Println("could not serialize proto message : " + name)
	}

	anything := &any.Any{
		TypeUrl: name,
		Value:   serialized,
	}

	baseMessage := BaseMessage.BaseMessage{Id: messageIndex, Message: anything}
	bytes, seconderr := proto.Marshal(&baseMessage)

	if seconderr != nil {
		fmt.Println("Error serializing the base message")
	}

	fmt.Println(fmt.Sprintln("Write %d",

		len(bytes)))
	fmt.Println(fmt.Sprintln("Writing to ", index))
	connectionMap[index].con.Write(bytes)
}

func RemoveConnection(index int) {
	delete(connectionMap, index)
}
