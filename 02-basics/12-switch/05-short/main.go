package main

import "fmt"

func main() {
	switch num := 5; num {
	case 5:
		fmt.Println("equal to 5")
	default:
		fmt.Println("not equal to 5")
	}
}