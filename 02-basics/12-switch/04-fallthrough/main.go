package main

import "fmt"

func main() {
	hello := "hello"

	// Before, breaks after case matched
	switch hello {
	case "hello":
		fmt.Println("hello")
	case "hi":
		fmt.Println("hi")
	default:
		fmt.Println("hey")
	}

	// After, continues to other cases
	switch hello {
	case "hello":
		fmt.Println("hello")
		fallthrough
	case "hi":
		fmt.Println("hi")
		fallthrough
	default:
		fmt.Println("hey")
	}
}