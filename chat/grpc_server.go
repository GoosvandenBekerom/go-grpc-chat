package chat

import (
	"fmt"
	"github.com/GoosvandenBekerom/go-grpc-chat/pb"
	"io"
	"log"
)

type RpcServer struct {
	// A map containing chat-rooms and their respective history of messages
	rooms map[string][]*pb.Message
}

func (s *RpcServer) BroadcastMessage(server pb.Chat_BroadcastMessageServer) error {
	room := fmt.Sprintf("chatroom %d", len(s.rooms)+1)
	s.rooms[room] = make([]*pb.Message, 100)

	for {
		msg, err := server.Recv()
		if err == io.EOF {
			log.Println("client has closed connection")
			break
		}
		if err != nil {
			log.Println("unable to read message from client:", err)
			return err
		}
		log.Printf("Received message to broadcast from %s: %v", room, msg)
		// Append received message to chatroom history
		s.rooms[room] = append(s.rooms[room], msg)
		// Stream message back to chatroom to broadcast
		if err = server.Send(msg); err != nil {
			log.Println("unable to write message to client:", err)
			return err
		}
	}
	return nil
}

func NewRpcServer() *RpcServer {
	return &RpcServer{
		rooms: make(map[string][]*pb.Message),
	}
}
