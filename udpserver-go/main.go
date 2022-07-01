package main

import (
	"fmt"
	"net"
)

func main() {
	udpAddr := net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 5500}
	ln, err := net.ListenUDP("udp", &udpAddr)

	if err != nil {
		panic(err)
	}

	for {
		buf := make([]byte, 1024)
		oob := make([]byte, 1024)
		_, _, _, addr, err := ln.ReadMsgUDP(buf, oob)

		if err != nil {
			panic(err)
		}

		fmt.Println(fmt.Sprintf("Message received: %s, from address %s:%d", buf, addr.IP.String(), addr.Port))
	}
}
