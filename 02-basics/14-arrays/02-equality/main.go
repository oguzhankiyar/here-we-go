package main

import "fmt"

func main() {
	// True
	first := [...]int{ 1, 2 }
	second := [...]int{ 1, 2 }
	fmt.Printf("%v == %v => %v\n", first, second, first == second)

	// Compile error if different sizes
	// third := [3]int{ 1, 2 }
	// fmt.Printf("%v == %v => %v\n", first, third, first == third)

	// False if order is not matched
	fourth := [2]int{ 2, 1 }
	fmt.Printf("%v == %v => %v\n", first, fourth, first == fourth)
}