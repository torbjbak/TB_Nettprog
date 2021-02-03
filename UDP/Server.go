package main

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

var count = 0

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func calculate(number1, number2 float64, operator string) []byte {
	var result float64
	switch strings.ToLower(operator) {
	case "plus":
	case "+":
		result = number1 + number2
		break
	case "minus":
	case "-":
		result = number1 - number2
		break
	case "multiply":
	case "*":
		result = number1 * number2
		break
	case "divide":
	case "/":
		result = number1 / number2
		break
	}
	return []byte(strconv.FormatFloat(result, 'f', -1, 64))
}

func main() {
	s, err := net.ResolveUDPAddr("udp4", "localhost:8000")
	if err != nil {
		fmt.Println(err)
		return
	}

	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer connection.Close()
	buffer := make([]byte, 1024)
	rand.Seed(time.Now().Unix())

	for {
		n, addr, err := connection.ReadFromUDP(buffer)
		fmt.Print("-> ", string(buffer[0:n-1]))

		input := strings.TrimSpace(string(buffer[0:n]))

		if strings.EqualFold(input, "exit") {
			fmt.Println("Exiting UDP server!")
			return
		}

		var data []byte
		calc := strings.Fields(input)

		if len(calc) == 3 {
			n1, err1 := strconv.ParseFloat(calc[0], 64)
			n2, err2 := strconv.ParseFloat(calc[2], 64)

			if err1 == nil && err2 == nil {
				data = calculate(n1, n2, calc[1])
			} else {
				data = []byte("Arguments are not numbers!")
			}
		} else {
			data = []byte("Incorrect number of arguments!")
		}

		_, err = connection.WriteToUDP(data, addr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
