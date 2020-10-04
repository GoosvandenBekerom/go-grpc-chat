package main

import (
	"fmt"
	"github.com/GoosvandenBekerom/go-grpc-chat/chat"
	"github.com/GoosvandenBekerom/go-grpc-chat/config"
	"github.com/GoosvandenBekerom/go-grpc-chat/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.GrpcServerPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterChatServer(s, chat.NewRpcServer())
	log.Printf("Starting gRPC server on tcp port %d\n", config.GrpcServerPort)
	log.Fatal(s.Serve(lis))
}
