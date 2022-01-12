package main

import (
	"constraints"
	"fmt"
)

func main() {
	v := min[int](1, -2, 3, 0, -4)
	fmt.Println(v)
}

func min[T constraints.Signed](arr ...T) T {
	min := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}

	return min
}