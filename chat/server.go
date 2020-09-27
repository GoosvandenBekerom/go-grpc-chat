package chat

import (
	"context"
	"fmt"
	"time"
)

type Server struct {
	history []*Message
}

func (s Server) Send(_ context.Context, msg *Message) (*Message, error) {
	s.history = append(s.history, msg)
	printMessage(msg)
	return msg, nil
}

func printMessage(m *Message) {
	println(fmt.Sprintf("%s: %s", time.Unix(m.Timestamp, 0).Format("Jan _2 15:04:05"), m.Content))
}

func NewServer() Server {
	return Server{
		history: make([]*Message, 50),
	}
}
