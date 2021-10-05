package main

import (
	"fmt"
	"time"
)

// Firstly, the time variable will be created by using time package
// After that, the message will be appeared in console by using fmt package
func main() {
	now := time.Now()
	fmt.Println("Halo! Now:", now)
}