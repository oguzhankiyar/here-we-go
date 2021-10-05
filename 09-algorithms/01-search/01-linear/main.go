package main

import "fmt"

func main() {
	items := []int{1, 5, 7, 9}

	index := LinearSearch(items, 7)
	if index == -1 {
		fmt.Println("notfound")
	} else {
		fmt.Println("found:", index)
	}
}

func LinearSearch(items []int, item int) int {
	for i := 0; i < len(items); i++ {
		if items[i] == item {
			return i
		}
	}

	return -1
}