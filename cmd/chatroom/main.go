package main

import (
	"fmt"
	"github.com/GoosvandenBekerom/go-grpc-chat/chat"
	"github.com/GoosvandenBekerom/go-grpc-chat/config"
	"github.com/GoosvandenBekerom/go-grpc-chat/pb"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func serveFrontEnd(response http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.Error(response, "Not found", http.StatusNotFound)
		return
	}
	if request.Method != "GET" {
		http.Error(response, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(response, request, "static/index.html")
}

func main() {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", config.GrpcServerHost, config.GrpcServerPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect to grpc server: %v", err)
	}
	defer conn.Close()
	c := pb.NewChatClient(conn)
	room := chat.NewRoom(c)
	go room.Run()
	http.HandleFunc("/", serveFrontEnd)
	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		chat.ServeWebSocket(room, w, r)
	})
	log.Println("Starting chatroom on port ", config.WebSocketServerPort)
	err = http.ListenAndServe(fmt.Sprintf(":%d", config.WebSocketServerPort), nil)
	if err != nil {
		log.Fatal("Room failed to start: ", err)
	}
}
