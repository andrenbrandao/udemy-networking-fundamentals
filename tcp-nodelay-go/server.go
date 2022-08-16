package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	port := ":" + "8000"

	// Create a listening socket.
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		// Accept new connections.
		c, err := l.Accept()
		if err != nil {
			log.Println(err)
			return
		}

		// Process newly accepted connection.
		go handleConnection(c)
	}
}
func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())

	for {
		// Read what has been sent from the client.
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}

		cdata := strings.TrimSpace(netData)
		if cdata == "GOPHER" {
			c.Write([]byte("GopherAcademy Advent 2019!"))
		}

		if cdata == "EXIT" {
			break
		}
	}
	c.Close()
}
