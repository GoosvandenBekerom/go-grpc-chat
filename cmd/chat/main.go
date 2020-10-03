package main

import (
	"fmt"
	"github.com/GoosvandenBekerom/go-grpc-chat/chat"
	"github.com/GoosvandenBekerom/go-grpc-chat/config"
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
	// TODO: separate client and server to separate applications that communicate over gRPC
	server := chat.NewServer()
	go server.Run()
	http.HandleFunc("/", serveFrontEnd)
	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		chat.ServeWebSocket(server, w, r)
	})
	log.Println("Starting chat server on port ", config.ServerPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.ServerPort), nil)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
