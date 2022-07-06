package main

import (
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

	conn, err := ln.Accept()
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("TCP Handshake Successful with: %s", conn.RemoteAddr().String()))

	count := 0

	for {
		buf := make([]byte, 1024)
		conn.Read(buf)

		if strings.HasPrefix(string(buf), "exit") {
			fmt.Println("Closing connection...")
			conn.Close()
			break
		}

		fmt.Println(fmt.Sprintf("Message received: %s, from address %s", buf, conn.RemoteAddr().String()))
		conn.Write([]byte(fmt.Sprintf("Message received number %d\n", count)))
		count += 1
	}
}
