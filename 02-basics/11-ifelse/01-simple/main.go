package main

import "fmt"

func main() {
	// Only if
	if 1 == 1 {
		fmt.Println("1 is equal to 1")
	}

	// If and else
	if 1 == 2 {
		fmt.Println("1 is equal to 1")
	} else {
		fmt.Println("1 is not equal to 1")
	}

	// If, else if and else
	const hey = "hey"
	if hey == "hello" {
		fmt.Println("hello")
	} else if hey == "hi" {
		fmt.Println("hi")
	} else {
		fmt.Println("hey")
	}
}