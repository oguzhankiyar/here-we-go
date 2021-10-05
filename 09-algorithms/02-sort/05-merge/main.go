package main

import "fmt"

func main() {
	items := []int{9, 5, 7, 3}

	sorted := MergeSort(items)

	fmt.Println("sorted:", sorted)
}

func MergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}

	mid := len(items) / 2

	left := MergeSort(items[:mid])
	right := MergeSort(items[mid:])

	return Merge(left, right)
}

func Merge(left, right []int) []int {
	size := len(left) + len(right)
	i, j := 0, 0
	merged := make([]int, size)

	for k := 0; k < size; k++ {
		if i > len(left) - 1 && j <= len(right) - 1 {
			merged[k] = right[j]
			j++
		} else if j > len(right) - 1 && i <= len(left) - 1 {
			merged[k] = left[i]
			i++
		} else if left[i] < right[j] {
			merged[k] = left[i]
			i++
		} else {
			merged[k] = right[j]
			j++
		}
	}

	return merged
}