package sockets

import "net"

func connect() {
	raddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		panic("Uh oh", err)
	}

	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		panic("Failed to dial up to UDP")
	}

	buffer := make([]byte, maxBufferSize)
	n, addr, err := conn.ReadFrom(buffer)
	if err != nil {
		panic("Failed to read from UDP socket")
	}

	buffer = make([]byte, maxBufferSize)
	// send message
	_, err = conn.WriteTo(buffer[:n], addr)
	if err != nil {
		panic("Failed to send message")
	}

}
