package sockets

import "net"

func startTCPServer() {
	tcpAddr, err := net.ResolveIPAddr(resolver, serverAddr)
	if err != nil {
		panic("Failed to get address")
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		panic("Failed to listen")
	}
	conn, err := listener.Accept()
	if err != nil {
		panic("Failed to accept connection")
	}
	// writing
	if _, err := conn.Write({message}); err != nil {
		panic("Failed to write message")
	}
	// reading
	buf := make([]byte, 512)
	n, err := conn.Read(buf[0:])
	if err != nil {
		panic("Failed to write")
	}
}
