package chat

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/GoosvandenBekerom/go-grpc-chat/pb"
	"io"
	"log"
	"time"
)

type Room struct {
	rpcClient  pb.ChatClient
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
}

func toJson(m *pb.Message) []byte {
	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return b
}

func (r *Room) handleRpcMessages(bmc pb.Chat_BroadcastMessageClient) {
	for {
		m, err := bmc.Recv()
		if err == io.EOF {
			log.Println("server has closed connection")
			break
		}
		if err != nil {
			log.Println("unable to read message from server:", err)
			panic(err)
		}
		log.Println("Message received from gRPC server: ", m)
		payload := toJson(m)
		log.Println("Sending message over WebSocket: ", string(payload))
		for client := range r.clients {
			select {
			case client.send <- payload:
			default:
				// client's "send" channel is unable to receive messages, execute graceful cleanup for this client
				close(client.send)
				delete(r.clients, client)
			}
		}
	}
}

func (r *Room) Run() {
	bmc, err := r.rpcClient.BroadcastMessage(context.Background())
	if err != nil {
		panic(fmt.Sprintf("unable to open gRPC broadcast message stream: %v", err))
	}

	go r.handleRpcMessages(bmc)

	for {
		select {
		case client := <-r.register:
			r.clients[client] = true
		case client := <-r.unregister:
			if _, ok := r.clients[client]; ok {
				delete(r.clients, client)
				close(client.send)
			}
		case bytes := <-r.broadcast:
			log.Println("Message received on WebSocket: ", string(bytes))
			var msg pb.Message
			if err := json.Unmarshal(bytes, &msg); err != nil {
				panic(fmt.Sprintf("unable to parse message: %v", string(bytes)))
			}
			msg.Timestamp = time.Now().Unix()
			log.Println("Forwarding message to gRPC server")
			err = bmc.Send(&msg)
			if err != nil {
				panic("unable to send message to gRPC server")
			}
		}
	}
}

func NewRoom(client pb.ChatClient) *Room {
	return &Room{
		rpcClient:  client,
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
	}
}
