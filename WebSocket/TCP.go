package main

import (
	"fmt"
	"net"
	"strings"
)

func handleTCP(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	size, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}
	text := strings.TrimSpace(strings.TrimSuffix(string(buffer[:size]), "\n"))
	handleHTTP(conn, text)
}

func handleHTTP(conn net.Conn, text string) {
	if strings.HasPrefix(text, "GET / HTTP/1.1\r\n") {
		body :=
			"<!DOCTYPE html>\n" +
				"<html>\n" +
				"<head>\n" +
				"<meta charset=\"UTF-8\" />\n" +
				"</head>\n" +
				"<body>\n" +
				"WebSocket test page\n" +
				"<script>\n" +
				"let ws = new WebSocket('ws://localhost:3001');\n" +
				"ws.onmessage = event => alert('Message from server: ' + event.data);\n" +
				"ws.onopen = () => ws.send('hello');\n" +
				"</script>\n" +
				"</body>\n" +
				"</html>"

		headers := []string{
			"HTTP/1.1 200 OK",
			"Content-Type: text/html; charset=utf-8",
			fmt.Sprintf("Content-Length: %d", len(body)),
			"",
			"",
		}

		_, err1 := conn.Write([]byte(strings.Join(headers, "\r\n")))
		if err1 != nil {
			fmt.Println(err1)
		}
		_, err2 := conn.Write([]byte(body))
		if err2 != nil {
			fmt.Println(err2)
		}
	}
	_ = conn.Close()
}
