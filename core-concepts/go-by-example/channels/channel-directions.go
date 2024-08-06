package main

import "fmt"

// The ping function sends a message to the pings channel
func ping(pings chan<- string, msg string) {
	pings <- msg // Send the message to the pings channel
}

// The pong function receives a message from the pings channel and sends it to the pongs channel
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings // Receive a message from the pings channel
	pongs <- msg   // Send the message to the pongs channel
}

func main() {
	// Create channels for sending and receiving messages
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	// Send a message to the pings channel
	ping(pings, "message sent")

	// Pass the message from the pings channel to the pongs channel
	pong(pings, pongs)

	// Print the message received from the pongs channel
	fmt.Println(<-pongs)
}
