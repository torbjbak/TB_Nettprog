package main

import (
	"fmt"
	"net"
)

func listen(port string, handler func(net.Conn)) {
	listener, err1 := net.Listen("tcp", "localhost:"+port)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	defer listener.Close()

	for {
		conn, err2 := listener.Accept()
		if err2 != nil {
			fmt.Println(err2)
			return
		}
		defer conn.Close()
		go handler(conn)
	}
}

func main() {
	go listen("3001", handleWS)
	listen("3000", handleTCP)
}
