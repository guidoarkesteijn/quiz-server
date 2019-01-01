package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	BaseMessage "guido.arkesteijn/quiz-server/Data"
	Player "guido.arkesteijn/quiz-server/Data/Player"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/twinj/uuid"
)

type Connection struct {
	player Player.Player
	Index  int
	active bool
	con    net.Conn
}

var connections []Connection
var i = 0
var messageCount int32 = 1

var stringToInt map[string]int

func HandleConnection(c net.Conn) {
	player := Player.Player{Guid: uuid.NewV4().String()}
	connection := Connection{player, i, true, c}
	connections = append(connections, connection)
	go Read(i, connection)

	fmt.Println("Added connection number:" + fmt.Sprintf("%d", connection.Index))
	i++
}

func Read(index int, connection Connection) {
	c := connection.con

	fmt.Printf("Serving %s\n", c.RemoteAddr().String())

	for {
		fmt.Println("start reading welcome message.")
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println("reading error " + err.Error())
			return
		}

		temp := strings.TrimSpace(string(netData))
		fmt.Println("Received: " + temp)
		if temp == "STOP" {
			break
		}

		//Only write to the client that is connecting.
		WriteSingle(index, &Player.PlayerJoined{Id: int32(index), Player: &connection.player})
	}

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

	typeInt := GetIntForString(name)
	baseMessage := BaseMessage.BaseMessage{Id: messageIndex, Message: anything, Type: typeInt}
	bytes, seconderr := proto.Marshal(&baseMessage)

	if seconderr != nil {
		fmt.Println("Error serializing the base message")
	}

	fmt.Println(fmt.Sprintln("Write %d",

		len(bytes)))
	fmt.Println(fmt.Sprintln("Writing to ", index))
	connections[index].con.Write(bytes)
}

func RemoveConnection(index int) {
	connections = RemoveAt(connections, index)
}

func RemoveAt(a []Connection, i int) []Connection {
	// Remove the element at index i from a.
	copy(a[i:], a[i+1:])                                      // Shift a[i+1:] left one index.
	a[len(a)-1] = Connection{Player.Player{}, -1, false, nil} // Erase last element (write zero value).
	a = a[:len(a)-1]                                          // Truncate slice.
	return a
}

func GetIntForString(t string) int32 {
	stringToInt = make(map[string]int)

	stringToInt["Data.Joined"] = 1

	fmt.Sprintln("TEST ", stringToInt[t])

	return int32(stringToInt[t])
}
