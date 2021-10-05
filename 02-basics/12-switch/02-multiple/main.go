package main

import "fmt"

func main() {
	message := "hello"

	switch message {
	case "hello", "hi":
		fmt.Println("greeting")
	default:
		fmt.Println("other")
	}
}