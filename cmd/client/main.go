package main

import (
	"context"
	"fmt"
	"github.com/GoosvandenBekerom/go-grpc-chat/chat"
	"github.com/GoosvandenBekerom/go-grpc-chat/config"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", config.ServerHost, config.ServerPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect to grpc server: %v", err)
	}
	defer conn.Close()
	client := chat.NewChatClient(conn)
	msg, err := client.Send(context.Background(), &chat.Message{
		Timestamp: time.Now().Unix(),
		Content:   "Hello, this message is sent from a client!",
	})
	if err != nil {
		log.Printf("Unable to send message: %v", err)
	} else {
		log.Printf("Message sent: %s", msg.Content)
	}
}
