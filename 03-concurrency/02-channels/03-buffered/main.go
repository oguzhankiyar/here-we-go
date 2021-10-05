package main

import (
	"fmt"
	"time"
)

func main() {
	// Scenario 1
	ch1 := make(chan string, 2)

	ch1 <- "Hello"
	ch1 <- "Hi"
	// When send a new message, the channel blocks until empty
	// ch1 <- "Hey"

	fmt.Println(<-ch1)
	fmt.Println(<-ch1)

	// Scenario 2

	write := func(ch chan int) {
		for i := 1; i <= 4; i++ {
			ch <- i
			fmt.Println("Sent:", i)
		}
		close(ch)
	}

	ch2 := make(chan int, 2)

	go write(ch2)
	time.Sleep(2 * time.Second)
	for v := range ch2 {
		fmt.Println("Received:", v)
		time.Sleep(2 * time.Second)
	}
}