package main

import "fmt"

func main() {
	Sample("Errorf", Errorf)
}

func Errorf() {
	// Creates an error with given message format

	err := fmt.Errorf("user is not found with id: %d", 100)
	fmt.Println(err)
}

func Sample(name string, fn func())  {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}