package main

import "fmt"

func main() {
	// Simple
	var year int = 2000
	year = 2000 + 20

	// You can use := form if one of them is undeclared before
	var month = 5
	month, day := 6, 20

	fmt.Println(day, month, year)

	// Swap
	var first, second = 1, 2
	first, second = second, first

	fmt.Println(first, second)
}
