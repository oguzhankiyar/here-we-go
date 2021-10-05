package main

import "fmt"

func main() {
	items := []int{9, 5, 7, 3}

	sorted := ShellSort(items)

	fmt.Println("sorted:", sorted)
}

func ShellSort(items []int) []int {
	length := len(items)

	for gap := length / 2; gap > 0; gap /= 2 {
		for i := gap; i < length; i++ {
			item := items[i]

			j := i
			for ; j >= gap && items[j - gap] > item; j -= gap {
				items[j] = items[j - gap]
			}

			items[j] = item
		}
	}

	return items
}