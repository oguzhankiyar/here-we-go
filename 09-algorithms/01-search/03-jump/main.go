package main

import (
	"fmt"
	"math"
)

func main() {
	items := []int{1, 5, 7, 9}

	index := JumpSearch(items, 7)
	if index == -1 {
		fmt.Println("notfound")
	} else {
		fmt.Println("found:", index)
	}
}

func JumpSearch(items []int, item int) int {
	itemSize := len(items)
	blockSize := int(math.Sqrt(float64(itemSize)))

	start := 0

	for {
		if items[start] >= item {
			break
		}

		if start > itemSize {
			break
		}

		start += blockSize
	}

	for i := start; i > 0; i-- {
		if items[i] == item {
			return i
		}
	}

	return -1
}