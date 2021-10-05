package main

import "fmt"

func main() {
	items := []int{9, 5, 7, 3}

	sorted := InsertionSort(items)

	fmt.Println("sorted:", sorted)
}

func InsertionSort(items []int) []int {
	length := len(items)

	for i := 1; i < length; i++ {
		j := i

		for j > 0 {
			if items[j - 1] > items[j] {
				Swap(items, j - 1, j)
			}

			j--
		}
	}

	return items
}

func Swap(items []int, i, j int) {
	items[i], items[j] = items[j], items[i]
}