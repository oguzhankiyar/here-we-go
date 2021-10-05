package main

import "fmt"

func main() {
	a, b := 2, 1

	if a % 2 == 0 {
		fmt.Println("a is even")

		if b % 2 != 0 {
			fmt.Println("b is odd")
		} else {
			fmt.Println("b is even")
		}
	}
}