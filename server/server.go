package server

import (
	"fmt"
	"net"
)

var connectionMap map[int]Connection

func StartServer(port int) {
	fmt.Println("Starting server")
	connectionMap = make(map[int]Connection)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer l.Close()

	fmt.Println("Waiting for clients")

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		HandleConnection(c)
	}
}
