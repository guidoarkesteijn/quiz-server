package server

import (
	"fmt"
	"net"
)

var letters = []string{"a", "b", "c", "d", "etc"}

func StartServer(port int) {
	fmt.Println("Starting server")

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
