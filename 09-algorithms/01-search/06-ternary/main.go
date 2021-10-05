package main

import "fmt"

func main() {
	items := []int{1, 5, 7, 9}

	index := TernarySearch(items, 7, 0, len(items) - 1)
	if index == -1 {
		fmt.Println("notfound")
	} else {
		fmt.Println("found:", index)
	}
}

func TernarySearch(items []int, item int, startIndex int, endIndex int) int {
	if startIndex <= endIndex {
		midFirst := startIndex + (endIndex - startIndex) / 3
		midSecond := midFirst + (endIndex - startIndex) / 3

		if items[midFirst] == item {
			return midFirst
		}

		if items[midSecond] == item {
			return midSecond
		}

		if item < items[midFirst] {
			return TernarySearch(items, item, startIndex, midFirst-1)
		}

		if item > items[midSecond] {
			return TernarySearch(items, item, midSecond + 1, endIndex)
		}

		return TernarySearch(items, item, midFirst + 1, midSecond - 1)
	}

	return -1
}