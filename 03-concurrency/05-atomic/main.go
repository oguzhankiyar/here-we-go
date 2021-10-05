package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var count int64
	ch := make(chan bool)

	for i := 0; i < 10; i++ {
		go func() {
			atomic.AddInt64(&count, 1)
			ch <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-ch
	}

	fmt.Println("count:", count)
}