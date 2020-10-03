package chat

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Server struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
	history    []*Message
}

func (s *Server) Run() {
	for {
		select {
		case client := <-s.register:
			s.clients[client] = true
		case client := <-s.unregister:
			if _, ok := s.clients[client]; ok {
				delete(s.clients, client)
				close(client.send)
			}
		case message := <-s.broadcast:
			log.Println("Message received: ", string(message))
			for client := range s.clients {
				select {
				case client.send <- message:
				default:
					// client's "send" channel is unable to receive messages, execute graceful cleanup for this client
					close(client.send)
					delete(s.clients, client)
				}
			}
		}
	}
}

func (s Server) Send(_ context.Context, msg *Message) (*Message, error) {
	s.history = append(s.history, msg)
	printMessage(msg)
	return msg, nil
}

func printMessage(m *Message) {
	println(fmt.Sprintf("[%s] %s: %s", time.Unix(m.Timestamp, 0).Format("Jan _2 15:04:05"), m.Username, m.Content))
}

func NewServer() *Server {
	return &Server{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
		history:    make([]*Message, 50),
	}
}
