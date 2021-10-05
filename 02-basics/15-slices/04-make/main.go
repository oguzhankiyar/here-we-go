package main

import "fmt"

func main() {
	slice := make([]int, 0, 10)
	fmt.Printf("slice = %v\n", slice)
	fmt.Printf("len(slice) = %v\n", len(slice))
	fmt.Printf("cap(slice) = %v\n", cap(slice))
}