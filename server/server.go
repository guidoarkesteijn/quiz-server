package server

import (
	"fmt"
	"net"
)

var connectionMap map[int]Connection

type ServerService struct {
	connected         bool
	onMessageReceived chan int
}

//New create new server service with onMessageReceived channel
func New() ServerService {
	return ServerService{onMessageReceived: make(chan int, 10)}
}

//Start start server on the given port.
func (s *ServerService) Start(port int) {
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
		go s.WaitForRead()
		HandleConnection(s, c)
	}
}

//WaitForRead test function to see how channels work.
func (s *ServerService) WaitForRead() {
	fmt.Println("Waiting for value")
	value := <-s.onMessageReceived

	fmt.Println("Value: ", value)
}
