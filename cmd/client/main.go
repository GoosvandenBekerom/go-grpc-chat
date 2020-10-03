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
	print("Choose a username: ")
	var user string
	_, err := fmt.Scanln(&user)
	check(err)
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", config.ServerHost, config.ServerPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect to grpc server: %v", err)
	}
	defer func() { check(conn.Close()) }()

	c := chat.NewChatClient(conn)

	for true {
		print("Send message ('q' to quit): ")
		var input string
		_, err := fmt.Scanln(&input)
		check(err)
		if input == "q" {
			break
		}
		sendMessage(c, user, input)
	}
}

func sendMessage(c chat.ChatClient, username, content string) {
	_, err := c.Send(context.Background(), &chat.Message{
		Timestamp: time.Now().Unix(),
		Username:  username,
		Content:   content,
	})
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
