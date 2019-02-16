package server

import (
	"bufio"
	"bytes"
	"fmt"
	"net"

	"github.com/project-quiz/quiz-go-model/message"

	"github.com/project-quiz/quiz-server/interpreter"
	"github.com/project-quiz/quiz-server/model"
	"github.com/twinj/uuid"
)

var messageCount int32 = 1

//HandleConnection handle incoming connection.
func HandleConnection(server *Service, c net.Conn) {
	connection := model.PlayerClient{Guid: uuid.NewV4().String(), NickName: "<UNKNOWN>", Con: c}

	playerClientList[connection.Guid] = connection
	go Read(server, &connection)

	fmt.Println("Added connection number:" + fmt.Sprintf("%d", len(playerClientList)))
}

func ToPlayerClient(m *message.Player) model.PlayerClient {
	return playerClientList[m.Guid]
}

//Read reading from the connection.
func Read(server *Service, connection *model.PlayerClient) {
	guid := connection.Guid
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
			server.onMessageReceived <- model.Result{Message: &m, PlayerClient: connection}
		}

		fmt.Println("Waiting for next message")
	}

	fmt.Println("Close: ", c.RemoteAddr().String())
	DeleteConnection(guid)
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

//DeleteConnection delete the given index from the connection map.
func DeleteConnection(guid string) {
	delete(playerClientList, guid)
}
