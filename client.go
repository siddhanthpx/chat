package main

import (
	"github.com/gorilla/websocket"
)

// client represents a single chatting user
type client struct {
	// socket is the websocket for this client
	socket *websocket.Conn
	// send is a channel on which messages are sent
	send chan []byte
	// room is the room that the client is chatting on
	room *room
}

func (c *client) read() {
	defer c.socket.Close()

	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}

		c.room.forward <- msg
	}
}

func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			return
		}
	}
}
