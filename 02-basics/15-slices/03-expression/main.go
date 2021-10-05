package main

import "fmt"

func main() {
	slice := []int{ 1, 2, 3, 4, 5 }

	fmt.Printf("slice[0:0] = %v\n", slice[0:0])
	fmt.Printf("slice[0:1] = %v\n", slice[0:1])
	fmt.Printf("slice[0:2] = %v\n", slice[0:2])
	fmt.Printf("slice[0:3] = %v\n", slice[0:3])
	fmt.Printf("slice[0:4] = %v\n", slice[0:4])

	fmt.Println()

	fmt.Printf("slice[:] = %v\n", slice[:])
	fmt.Printf("slice[0:] = %v\n", slice[0:])
	fmt.Printf("slice[:3] = %v\n", slice[:3])
}