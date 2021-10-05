package main

import "fmt"

func main() {
	var items = [2]int{ 1, 2 }
	for _, v := range items {
		fmt.Printf("%d ", v)
	}

	fmt.Println()

	items[0] = 3
	for _, v := range items {
		fmt.Printf("%d ", v)
	}
}