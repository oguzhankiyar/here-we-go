package main

import (
	"fmt"
	"math"
)

func main() {
	items := []int{1, 5, 7, 9}

	index := ExponentialSearch(items, 7, 0, len(items) - 1)
	if index == -1 {
		fmt.Println("notfound")
	} else {
		fmt.Println("found:", index)
	}
}

func BinarySearch(items []int, item int, startIndex int, endIndex int) int {
	if startIndex > endIndex || len(items) == 0 {
		return -1
	}

	middleIndex := (startIndex + endIndex) / 2

	if items[middleIndex] < item {
		return BinarySearch(items, item, middleIndex + 1, endIndex)
	} else if items[middleIndex] > item {
		return BinarySearch(items, item, startIndex, middleIndex - 1)
	} else {
		return middleIndex
	}
}

func ExponentialSearch(items []int, item int, startIndex int, endIndex int) int {
	if startIndex > endIndex || len(items) == 0 {
		return -1
	}

	i := int(math.Pow(2, 0))

	for i < endIndex - startIndex {
		if items[i] < item {
			i *= 2
		} else {
			break
		}
	}

	return BinarySearch(items, item, i / 2, i)
}