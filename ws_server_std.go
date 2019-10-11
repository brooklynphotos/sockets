package sockets

import "net/http"

func startWSServer() {
	mux := http.NewServeMux()
	mux.Handle("/", websocket.Handle(func(conn *websocket.Conn) {
		func(){
			for{
				message := messageType{}
				if err := websocket.JSON.Receive(conn, &message); err != nil {
						// handle error
				}
						.......     
				// send message
				if err := websocket.JSON.Send(conn, message); err != nil {
						// handle error
				}    
			}
		}
	}))
}
