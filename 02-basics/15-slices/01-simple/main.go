package main

import "fmt"

func main() {
	// Array
	var array [3]int
	fmt.Println("array:", array)
	fmt.Println("len(array):", len(array))
	fmt.Println("cap(array):", cap(array))

	fmt.Println()

	// Slice
	var slice []int
	fmt.Println("slice:", slice)
	fmt.Println("len(slice):", len(slice))
	fmt.Println("cap(slice):", cap(slice))
	if slice == nil {
		fmt.Println("slice is nil!")
	}
}