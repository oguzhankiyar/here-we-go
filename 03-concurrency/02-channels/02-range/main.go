package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 100)

	go calc(ch)

	// Waits until closed
	for i := range ch {
		fmt.Printf("%d ", i)
	}
}

func calc(ch chan int) {
	for	i := 1; i <= 10; i++ {
		ch <- i
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}