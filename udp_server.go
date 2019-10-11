package sockets

import "net"

func startUDPServer() {
	udpAddr, err := net.ResolveUDPAddr(resolver, serverAddr)
	if err != nil {
		panic("Failed to get address")
	}
	listener, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		panic("Failed to listen")
	}

	// send message
	buffer := make([]byte, maxBufferSize)
	_, err = conn.WriteToUDP(buffer[:n], addr)
	if err != nil {
		panic("Failed to send message")
	}

	buffer := make([]byte, maxBufferSize)
	n, addr, err := conn.ReadFromUDP(buffer)
	if err != nil {
		panic("Failed to read from UDP socket")
	}
}
