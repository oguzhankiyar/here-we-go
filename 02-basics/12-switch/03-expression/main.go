package main

import "fmt"

func main() {
	num := 5

	switch {
	case num < 3:
		fmt.Println("less than 3")
	case num > 3:
		fmt.Println("greater than 3")
	default:
		fmt.Println("equal to 3")
	}
}