package main

import "fmt"

func main() {
	var numbers [3]int

	numbers = [3]int{
		1,
		2,
		3,
	}
	fmt.Println(numbers)

	numbers = [3]int{
		1: 1,
		0: 2,
		2: 3,
	}
	fmt.Println(numbers)

	// Assigns zero value if not initialized
	numbers = [3]int{
		2: 99,
		1: 98,
	}
	fmt.Println(numbers)

	// Mix
	numbers = [3]int{
		1: 98,
		99,
	}
	fmt.Println(numbers)
}