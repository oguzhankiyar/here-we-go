package main

import (
	"fmt"
)

func main() {
	items := []int{9, 5, 7, 3}

	sorted := HeapSort(items)

	fmt.Println("sorted:", sorted)
}

func HeapSort(items []int) []int {
	length := len(items)

	for i := length / 2; i >= 0; i-- {
		HeapIt(items, i, length)
	}

	for i := length; i > 1; i-- {
		Swap(items, 0, i - 1)
		HeapIt(items, 0, i - 1)
	}

	return items
}

func HeapIt(items []int, root, length int) {
	max := root
	left := root * 2 + 1
	right := root * 2 + 2

	if left < length && items[left] > items[max] {
		max = left
	}

	if right < length && items[right] > items[max] {
		max = right
	}

	if max != root {
		Swap(items, root, max)
		HeapIt(items, max, length)
	}
}

func Swap(items []int, i, j int) {
	items[i], items[j] = items[j], items[i]
}