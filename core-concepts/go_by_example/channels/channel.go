package main

import (
	"fmt"
	"sync"
)

func grTest(from string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	var wg sync.WaitGroup

	messages := make(chan string)

	go func() { messages <- "ping" }()

	wg.Add(1)
	go grTest("goroutine", &wg)

	msg := <-messages
	fmt.Println(msg)

	wg.Wait()
}
