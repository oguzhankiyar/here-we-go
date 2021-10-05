package main

import "fmt"

func main() {
	var a = 100

	address := &a
	fmt.Printf("address = &a => %d\n", address)

	pointer := *address
	fmt.Printf("pointer = *address => %d\n", pointer)
}