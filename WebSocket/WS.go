package main

import (
	"bufio"
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"net"
	"strings"
)

func handleWS(conn net.Conn) {
	buf := bufio.NewReader(conn)
	getPrefix, err1 := buf.ReadString('\n')
	if err1 != nil {
		fmt.Println(err1)
	}

	if strings.HasPrefix(getPrefix, "GET / HTTP/1.1") {
		for {
			s, err2 := buf.ReadString('\n')
			if err2 != nil && err2 != io.EOF {
				fmt.Println(err2)
				break
			}

			text := strings.TrimSpace(strings.TrimSuffix(s, "\n"))
			if len(text) > 0 && text != "\n" {
				if strings.Contains(s, "Sec-WebSocket-Key:") {
					key := strings.TrimSpace(strings.Split(text, ":")[1])
					handshakeWS(conn, key)
					readWS(conn)
					break
				}
			}
		}
	}
	_ = conn.Close()
}

func handshakeWS(conn net.Conn, key string) {
	h := sha1.New()
	h.Write([]byte(key))
	h.Write([]byte("258EAFA5-E914-47DA-95CA-C5AB0DC85B11"))
	hash := base64.StdEncoding.EncodeToString(h.Sum(nil))

	headers := []string{
		"HTTP/1.1 101 Switching Protocols",
		"Upgrade: websocket",
		"Connection: Upgrade",
		"Sec-WebSocket-Accept: " + hash,
		"Sec-WebSocket-Protocol: chat",
		"",
		"",
	}
	_, _ = conn.Write([]byte(strings.Join(headers, "\r\n")))
}

func readWS(conn net.Conn) {
	reader := bufio.NewReader(conn)
	buf := make([]byte, 126)
	for {
		size, err := reader.Read(buf[0:])
		if err != nil {
			fmt.Println(err)
		}

		if buf[0] == '\x88' {
			break
		}

		maskStart := 2
		dataStart := maskStart + 4

		fmt.Print("Receiving from client: ")
		for i := dataStart; i < size; i++ {
			b := buf[i] ^ buf[maskStart+((i-dataStart)%4)]
			fmt.Printf("%s", string(b))
		}
		fmt.Printf("\n")

		writeWS(conn, []byte("Message?"))
	}
}

func writeWS(conn net.Conn, msg []byte) {
	if len(msg) > 126 {
		fmt.Println("Message too long!")
		return
	}

	var buffer bytes.Buffer
	buffer.WriteByte('\x81') // Hex 81 er desimal 129.
	buffer.WriteByte(byte(len(msg)))
	buffer.Write(msg)

	_, err := conn.Write(buffer.Bytes())
	if err != nil {
		fmt.Println(err)
	}
}
