package server

import (
	"bufio"
	"bytes"
	"fmt"
	"net"

	"github.com/golang/protobuf/proto"
	"github.com/project-quiz/quiz-go-model/model"
	"github.com/project-quiz/quiz-server/interpreter"
	"github.com/twinj/uuid"
)

//Connection this contains the data for an existing connection.
type Connection struct {
	player model.Player
	Index  int
	active bool
	con    net.Conn
}

var i = 0
var messageCount int32 = 1

//HandleConnection handle incoming connection.
func HandleConnection(server *Service, c net.Conn) {
	player := model.Player{Guid: uuid.NewV4().String(), Nickname: "<UNKNOWN>"}
	connection := Connection{player, i, true, c}

	connectionMap[i] = connection
	go Read(i, server, connection)

	fmt.Println("Added connection number:" + fmt.Sprintf("%d", connection.Index))
	i++
}

//Read reading from the connection.
func Read(index int, server *Service, connection Connection) {
	c := connection.con

	fmt.Printf("Serving %s\n", c.RemoteAddr().String())

	reader := bufio.NewReader(c)
	scanner := bufio.NewScanner(reader)
	scanner.Split(scanCRLF)

	for scanner.Scan() {
		bytes := scanner.Bytes()
		m, err := interpreter.ReadBaseMessage(bytes)

		if err != nil {
			fmt.Println("Error reading base message: ", err.Error())
			break
		} else {
			//send combined Result message (so the connection is linked to the message)
			server.onMessageReceived <- Result{Message: m, Connection: connection}
		}

		fmt.Println("Waiting for next message")

		/*
			playerJoin := model.PlayerJoin{}
			err2 := ptypes.UnmarshalAny(m.Message, &playerJoin)

			if err != nil || err2 != nil {
				fmt.Println("Read base message ", err2.Error())
				break
			} else {
			}
				if err != nil || err2 != nil {
					fmt.Println(err2.Error())
					break
				} else {
					p := &connection.player
					p.Nickname = playerJoin.Nickname
					//Only write to the client that is connecting.
					WriteSingle(index, &model.PlayerJoined{Guid: connection.player.Guid, Player: &connection.player})
				}
				server.onMessageReceived <- m
		*/
		//c.Server.onNewMessage(c, strings.ToUpper(hex.EncodeToString(scanner.Bytes())+"0d0a"))
	}

	fmt.Println("Close: ", c.RemoteAddr().String())
	DeleteConnection(index)
	c.Close()
}

func scanCRLF(data []byte, atEOF bool) (advance int, token []byte, err error) {
	skipBytes := []byte{'[', 'E', 'N', 'D', ']'}

	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, skipBytes); i >= 0 {
		// We have a full newline-terminated line.
		return i + len(skipBytes), data[0:i], nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
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
	bytes, err := interpreter.WriteBaseMessage(message)

	if err != nil {
		fmt.Println("WriteWithIndex: ", err.Error())
	} else {
		//write to the correct connection.
		connectionMap[index].con.Write(bytes)
	}
}

//DeleteConnection delete the given index from the connection map.
func DeleteConnection(index int) {
	delete(connectionMap, index)
}
