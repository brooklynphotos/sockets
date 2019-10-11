package sockets

import "net"

func connectTCPClient() {
	tcpAddr, err := net.ResolveTCPAddr(resolver, serverAddr)
	if err != nil {
		panic("Uh oh", err)
	}

	conn, err := net.DialTCP(network, nil, tcpAddr)

	// send message
	_, err = conn.Write({message})
	if err != nil {
		panic("Failed to send message")
	}

	var buf [{buffSize}]byte
	_, err = conn.Read(buf[0:])
	if err != nil {
		panic("Failed to read message")
	}
}
