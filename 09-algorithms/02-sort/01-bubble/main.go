package main

import "fmt"

func main() {
	items := []int{9, 5, 7, 3}

	sorted := BubbleSort(items)

	fmt.Println("sorted:", sorted)
}

func BubbleSort(items []int) []int {
	for i := 0; i < len(items); i++ {
		swapped := false

		for j := 0; j < len(items) - i - 1; j++ {
			if items[j] > items[j + 1] {
				items[j], items[j + 1] = items[j + 1], items[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}

	return items
}