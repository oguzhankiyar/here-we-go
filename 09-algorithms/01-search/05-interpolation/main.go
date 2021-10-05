package main

import "fmt"

func main() {
	items := []int{1, 5, 7, 9}

	index := InterpolationSearch(items, 7, 0, len(items) - 1)
	if index == -1 {
		fmt.Println("notfound")
	} else {
		fmt.Println("found:", index)
	}
}

func InterpolationSearch(items []int, item int, startIndex int, endIndex int) int {
	var dist, valRange, indexRange, estimate int
	var fraction float64

	for startIndex <= endIndex && item >= items[startIndex] && item <= items[endIndex] {
		dist = item - items[startIndex]
		valRange = items[endIndex] - items[startIndex]
		fraction = float64(dist) / float64(valRange)
		indexRange = endIndex - startIndex
		estimate = int(float64(startIndex) + (fraction * float64(indexRange)))

		if items[estimate] == item {
			return estimate
		}

		if items[estimate] < item {
			startIndex = estimate + 1
		} else {
			endIndex = estimate - 1
		}
	}
	return -1
}
