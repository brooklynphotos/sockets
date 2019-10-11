package sockets

import "golang.org/x/net/websocket"

func connectWSClient() {
	conn, err := websocket.Dial("{schema}://{host}:{port}", "", op.Origin)
	if err != nil {
		panic("Failed to dial into websocket url")
	}
	defer conn.Close()

	// send message
	if err = websocket.JSON.Send(conn, {message}); err != nil {
		panic("Failed to send message")
	}

	// receive message
	message := messageType{}
	if err := websocket.JSON.Receive(conn, &message); err != nil {
		panic("Failed to receive message")
	}
}
