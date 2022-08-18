package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	udpAddr := net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 5500}
	ln, err := net.ListenUDP("udp4", &udpAddr)

	defer ln.Close()

	if err != nil {
		panic(err)
	}

	buf := make([]byte, 1024)
	count := 0

	for {
		n, addr, err := ln.ReadFromUDP(buf)

		if err != nil {
			panic(err)
		}

		if strings.TrimSpace(string(buf[:n-1])) == "exit" {
			fmt.Println("Exiting UDP server...")
			return
		}

		fmt.Println(fmt.Sprintf("Message received: %s, from address %s:%d", buf[:n-2], addr.IP.String(), addr.Port))

		data := []byte(fmt.Sprintf("Request number %d to this server", count))
		count += 1
		_, err = ln.WriteToUDP(data, addr)

		if err != nil {
			panic(err)
		}
	}
}
