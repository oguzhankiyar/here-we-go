package main

import "fmt"

func main() {
	hello := "hello"

	// In Go, break is not required for switch
	switch hello {
	case "hello":
		fmt.Println("hello")
	case "hi":
		fmt.Println("hi")
	default:
		fmt.Println("hey")
	}
}