package main

import (
	"fmt"
	"time"
)

// worker simulates a task that takes one second to complete and then signals completion.
func worker(done chan bool) {
	fmt.Print("working...") // Print that the worker is starting
	time.Sleep(time.Second) // Simulate a task by sleeping for 1 second
	fmt.Println("done")     // Print that the task is complete

	done <- true // Send a signal to the channel indicating completion
}

func main() {
	// Create a channel to signal that the worker has finished its task
	done := make(chan bool, 1)

	// Start the worker function as a goroutine
	go worker(done)

	// Wait for the worker to send a signal through the channel
	<-done
}
