package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	target := "localhost:8000"

	raddr, err := net.ResolveTCPAddr("tcp", target)
	if err != nil {
		log.Fatal(err)
	}

	// Establish a connection with the server.
	conn, err := net.DialTCP("tcp", nil, raddr)
	if err != nil {
		log.Fatal(err)
	}

	// conn.SetNoDelay(false) // Disable TCP_NODELAY; Nagle's Algorithm takes action.

	fmt.Println("Sending Gophers down the pipe...")

	for i := 0; i < 5; i++ {
		// Send the word "GOPHER" to the open connection.
		_, err = conn.Write([]byte("GOPHER\n"))
		if err != nil {
			log.Fatal(err)
		}
	}
}
