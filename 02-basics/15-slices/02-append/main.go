package main

import "fmt"

func main() {
	var slice []int
	fmt.Println("slice:", slice)

	slice = append(slice, 1)
	fmt.Println("slice:", slice)

	slice = append(slice, 2, 3, 4, 5)
	fmt.Println("slice:", slice)

	var other = []int{ 7, 8 }
	slice = append(slice, other...)
	fmt.Println("slice:", slice)
}