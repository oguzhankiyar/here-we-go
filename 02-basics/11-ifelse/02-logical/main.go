package main

import "fmt"

func main() {
	const a = 1
	const b = 2

	if a % 2 == 0 && b % 2 == 0 {
		fmt.Println("a and b are even")
	}

	if a % 2 == 0 || b % 2 == 0 {
		fmt.Println("a or b are even")
	}

	if !(a % 2 == 0) {
		fmt.Println("a is not even")
	}
}