package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		time.Sleep(4 * time.Second)
		ch <- 1
	}()

	select {
	case x := <-ch:
		fmt.Println("Received", x)
	case <-time.After(5 * time.Second):
		fmt.Println("Timeout")
	}
}