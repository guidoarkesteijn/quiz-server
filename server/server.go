package server

import (
	"fmt"
	"net"
)

var connectionMap map[int]Connection

//Service this service contains all code to start an tcp server on the given port create with server.New() and then call func start with port to start it.
type Service struct {
	connected         bool
	onMessageReceived chan int
}

//New create new server service with onMessageReceived channel
func New() Service {
	return Service{onMessageReceived: make(chan int, 10)}
}

//Start start server on the given port.
func (s *Service) Start(port int) {
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
	fmt.Println("Waiting for value")
	value := <-s.onMessageReceived

	fmt.Println("Value: ", value)
}
