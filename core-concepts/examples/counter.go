package main

import (
	"fmt"
	"time"
)

func printNumber(channelTest chan int) {
	for i := 0; i < 10; i++ {
		channelTest <- i
		time.Sleep(time.Millisecond * 500)
	}
	close(channelTest)
}

func printNumberDelayed(channelTest2 chan int) {
	for i := 10; i > 0; i-- {
		channelTest2 <- i
		time.Sleep(time.Millisecond * 300)
	}
	close(channelTest2)
}

func main() {
	channelTest := make(chan int)
	go printNumber(channelTest)

	channelTest2 := make(chan int)
	go printNumberDelayed(channelTest2)

	// Control de canales cerrados
	channel1Closed := false
	channel2Closed := false

	for {
		// Si ambos canales están cerrados, se termina el bucle
		if channel1Closed && channel2Closed {
			break
		}

		select {
		case num, ok := <-channelTest:
			if ok {
				fmt.Println("Desde channelTest:", num)
			} else {
				channel1Closed = true // Marcar channelTest como cerrado
			}
		case num, ok := <-channelTest2:
			if ok {
				fmt.Println("Desde channelTest2:", num)
			} else {
				channel2Closed = true // Marcar channelTest2 como cerrado
			}
		}
	}
}
