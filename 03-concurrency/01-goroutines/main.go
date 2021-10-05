package main

import (
	"fmt"
	"time"
)

func main() {
	// Existing
	go printLoop()

	fmt.Println("Don't wait for it")

	// New
	go func() {
		fmt.Println("new func")
	}()

	// With Loop
	// Pass arguments into the function
	for i := 0; i < 3; i++ {
		go func(i int) {
			fmt.Printf("Run: %d\n", i)
		}(i)
	}

	time.Sleep(2 * time.Second)
}

func printLoop() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}
}