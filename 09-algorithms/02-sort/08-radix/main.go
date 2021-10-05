package main

import (
	"fmt"
)

func main() {
	items := []int{9, 5, 7, 3}

	sorted := RadixSort(items)

	fmt.Println("sorted:", sorted)
}

func RadixSort(items []int) []int {
	length := len(items)

	max := items[0]

	for i := 0; i < length; i++ {
		if items[i] > max {
			max = items[i]
		}
	}

	for i := 1; max / i > 0; i *= 10 {
		CountSort(items, i)
	}

	return items
}

func CountSort(items []int, exp int) {
	length := len(items)
	count := [10]int{0}
	output := make([]int, length)

	for i := 0; i < length; i++ {
		count[(items[i] / exp) % 10]++
	}

	for i := 1; i < 10; i++ {
		count[i] += count[i - 1]
	}

	for i := length - 1; i >= 0; i-- {
		output[count[(items[i] / exp) % 10] - 1] = items[i]
		count[(items[i] / exp) % 10]--
	}

	for i := 0; i < length; i++ {
		items[i] = output[i]
	}
}