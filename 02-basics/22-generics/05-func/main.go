package main

import (
	"fmt"
)

func main() {
	vi := []int{1, 2, 3, 4, 5, 6}
	vi = filter(vi, func(v int) bool {
		return v < 4
	})
	fmt.Println(vi)
}

func filter[T any](a []T, f func(T) bool) []T {
	var n []T
	for _, e := range a {
		if f(e) {
			n = append(n, e)
		}
	}
	return n
}