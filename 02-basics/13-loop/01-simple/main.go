package main

import "fmt"

func main() {
	// Type 1
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i + 1)
	}

	fmt.Println()

	// Type 2
	j := 0
	for ; j < 10 ; {
		fmt.Printf("%d ", j + 1)
		j++
	}
}