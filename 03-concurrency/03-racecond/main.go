package main

import (
	"fmt"
)

var counter = 0

func main() {
	ch := make(chan int, 1)

	total := 10

	for i := 0; i < total; i++ {
		go func() {
			ch <- 1
			v := counter
			<- ch
			v++
			counter = v
		}()
	}
	fmt.Println("Counter: ", counter)
}