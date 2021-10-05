package main

import "fmt"

// Firstly, this method will execute
func init() {
	fmt.Println("Init!")
}

// After init, this method will execute
func main() {
	fmt.Println("Main!")
}