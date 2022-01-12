package main

import "fmt"

func main() {
	fmt.Println(find([]int{1, 2, 3, 4, 5, 6}, 5))
}

func find[T comparable](a []T, v T) int {
	for i, e := range a {
		if e == v {
			return i
		}
	}
	return -1
}