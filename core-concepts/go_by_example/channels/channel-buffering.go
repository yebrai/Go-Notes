package main

import "fmt"

func main() {
	// Create a buffered channel with a capacity of 2
	messages := make(chan string, 2)

	// Send "buffered" into the channel
	messages <- "buffered"

	// Launch a goroutine to send "channel" into the channel
	go func() {
		messages <- "channel"
	}()

	// Receive and print messages from the channel
	fmt.Println(<-messages) // Receives "buffered"
	fmt.Println(<-messages) // Receives "channel"
	fmt.Println("test")
}
