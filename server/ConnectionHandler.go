package server

import (
	"bufio"
	"bytes"
	"fmt"
	"net"

	"github.com/golang/protobuf/proto"
	"github.com/project-quiz/quiz-go-model/model"
	"github.com/project-quiz/quiz-server/interpreter"
	"github.com/project-quiz/quiz-server/message"
	"github.com/twinj/uuid"
)

var i = 0
var messageCount int32 = 1

//HandleConnection handle incoming connection.
func HandleConnection(server *Service, c net.Conn) {
	player := model.Player{Guid: uuid.NewV4().String(), Nickname: "<UNKNOWN>"}
	connection := message.Connection{Player: &player, Index: i, Con: c}

	connectionMap[i] = connection
	go Read(i, server, &connection)

	fmt.Println("Added connection number:" + fmt.Sprintf("%d", connection.Index))
	i++
}

//Read reading from the connection.
func Read(index int, server *Service, connection *message.Connection) {
	c := connection.Con

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
			server.onMessageReceived <- message.Result{Message: &m, Connection: connection}
		}

		fmt.Println("Waiting for next message")
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
		connectionMap[index].Con.Write(bytes)
	}
}

//DeleteConnection delete the given index from the connection map.
func DeleteConnection(index int) {
	delete(connectionMap, index)
}
