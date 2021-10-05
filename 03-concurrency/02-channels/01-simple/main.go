package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		ch <- -1
	}()

	v := <- ch
	fmt.Println("Received", v)
}