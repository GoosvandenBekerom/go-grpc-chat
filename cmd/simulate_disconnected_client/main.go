package main

import "fmt"

// I wrote this to get a grasp of what happens in a select statement when a case is to send something to a channel
// that is not ready, in my head its a kind of representation of a chat client being unable to listen to messages
// the default case in the bottom will only be triggered when the "send" channel is unable to receive a message
// so we can use it to gracefully cleanup that specific client / channel
func main() {
	send := make(chan int)
	broadcast := make(chan int)
	defer close(broadcast)

	for i := 0; i < 10; i++ {
		go func() { broadcast <- i }()
		go func() {
			if i < 9 {
				fmt.Println(<-send)
			}
		}()
		select {
		case msg := <-broadcast:
			fmt.Println("broadcast received", msg)

			select {
			case send <- msg:
				fmt.Println("sent", msg, "to 'send' channel")
			default:
				close(send)
				fmt.Println("'send' channel closed")
			}
		}
	}
}
