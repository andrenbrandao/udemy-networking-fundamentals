package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	tcpAddr := net.TCPAddr{IP: net.ParseIP("127.0.0.1"), Port: 5500}
	ln, err := net.ListenTCP("tcp", &tcpAddr)

	if err != nil {
		panic(err)
	}

	count := 0

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		count++

		go handleConnection(conn, count)
	}
}

func handleConnection(conn net.Conn, count int) {
	fmt.Println(fmt.Sprintf("TCP Handshake Successful with: %s", conn.RemoteAddr().String()))
	conn.Write([]byte(fmt.Sprintf("Connection number %d\n", count)))

	msgCount := 1

	for {
		s, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			panic(err)
		}

		if strings.TrimSpace(s) == "exit" {
			fmt.Println(fmt.Sprintf("Closing connection number %d", count))
			conn.Close()
			return
		}

		fmt.Println(fmt.Sprintf("Message received: '%s', from address %s", strings.TrimSuffix(s, "\n"), conn.RemoteAddr().String()))
		conn.Write([]byte(fmt.Sprintf("Message received number %d\n", msgCount)))
		msgCount++
	}
}
