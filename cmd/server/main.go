package main

import (
	"fmt"
	"github.com/GoosvandenBekerom/go-grpc-chat/chat"
	"github.com/GoosvandenBekerom/go-grpc-chat/config"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.ServerPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	chat.RegisterChatServer(s, chat.NewServer())
	log.Printf("Starting gRPC server on tcp port %d", config.ServerPort)
	log.Fatal(s.Serve(lis))
}
