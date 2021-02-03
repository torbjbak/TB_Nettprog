package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	s, err := net.ResolveUDPAddr("udp4", "localhost:8000")
	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("The UDP server is %s\n", c.RemoteAddr().String())
	defer c.Close()

	for {
		fmt.Println("Enter your math expression or type 'exit' to exit!\n" +
			"Use format 'number1 operator number2'")
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		data := []byte(text + "\n")
		_, err = c.Write(data)
		if strings.EqualFold(strings.TrimSpace(string(data)), "exit") {
			fmt.Println("Exiting UDP client!")
			return
		}

		if err != nil {
			fmt.Println(err)
			return
		}

		buffer := make([]byte, 1024)
		n, _, err := c.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Result: %s\n\n", string(buffer[0:n]))
	}
}
