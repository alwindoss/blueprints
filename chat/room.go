package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type room struct {
	forward chan []byte
	join    chan *client
	leave   chan *client
	clients map[*client]bool
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			// joining
			r.clients[client] = true
		case client := <-r.leave:
			// leaving
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.forward:
			// forward message to all clients
			for client := range r.clients {
				client.send <- msg
			}
		}
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  socketBufferSize,
	WriteBufferSize: socketBufferSize,
}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Println("Client Connected")
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}
	client := &client{
		room:   r,
		send:   make(chan []byte, messageBufferSize),
		socket: socket,
	}
	r.join <- client
	defer func() {
		r.leave <- client
	}()
	go client.write()
	client.read()
}

func newRoom() *room {
	return &room{
		clients: make(map[*client]bool),
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
	}
}
