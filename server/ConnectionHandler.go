package server

import (
	"bufio"
	"bytes"
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

func HandleConnection(server *ServerService, c net.Conn) {
	player := Player.Player{Guid: uuid.NewV4().String(), Nickname: "<UNKNOWN>"}
	connection := Connection{player, i, true, c}

	connectionMap[i] = connection
	go Read(i, server, connection)

	fmt.Println("Added connection number:" + fmt.Sprintf("%d", connection.Index))
	i++
}

func WaitForRead(message chan int) {
	fmt.Println("Waiting for message value.")

	value := <-message

	fmt.Println("Message value: ", value)
}

func Read(index int, server *ServerService, connection Connection) {
	c := connection.con

	fmt.Printf("Serving %s\n", c.RemoteAddr().String())

	reader := bufio.NewReader(c)
	scanner := bufio.NewScanner(reader)
	scanner.Split(ScanCRLF)

	for scanner.Scan() {
		bytes := scanner.Bytes()

		var m BaseMessage.BaseMessage
		err := proto.Unmarshal(bytes, &m)

		if err != nil {
			fmt.Println("Unmarshal Error: ", err.Error())
			break
		}

		fmt.Println(m.Message.TypeUrl)
		fmt.Println(proto.MessageType(m.Message.TypeUrl))

		playerJoin := Player.PlayerJoin{}
		err2 := ptypes.UnmarshalAny(m.Message, &playerJoin)

		if err != nil || err2 != nil {
			fmt.Println(err2.Error())
			break
		} else {
			p := &connection.player
			p.Nickname = playerJoin.Nickname

			server.onMessageReceived <- 5
			//Only write to the client that is connecting.
			WriteSingle(index, &Welcome.Welcome{Player: &connection.player})
		}
		//c.Server.onNewMessage(c, strings.ToUpper(hex.EncodeToString(scanner.Bytes())+"0d0a"))
	}

	fmt.Println("Close: ", c.RemoteAddr().String())
	RemoveConnection(index)
	c.Close()
}

func ScanCRLF(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, []byte{'\r', '\n'}); i >= 0 {
		// We have a full newline-terminated line.
		return i + 2, dropCR(data[0:i]), nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), dropCR(data), nil
	}
	// Request more data.
	return 0, nil, nil
}

// dropCR drops a terminal \r from the data.
func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
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

	fmt.Println(fmt.Sprintln("Write ", len(bytes)))
	fmt.Println(fmt.Sprintln("Writing to ", index))
	connectionMap[index].con.Write(bytes)
}

func RemoveConnection(index int) {
	delete(connectionMap, index)
}
