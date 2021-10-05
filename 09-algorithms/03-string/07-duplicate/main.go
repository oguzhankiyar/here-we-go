package main

import "fmt"

func main() {
	fn := func(str string) {
		result := Duplicate(str)
		fmt.Printf("%q -> %q\n", str, result)
	}

	fn("gopher is learning golang!")
}

func Duplicate(str string) []rune {
	counts := make(map[rune]int)

	for _, c := range str {
		if _, ok := counts[c]; ok {
			counts[c]++
		} else {
			counts[c] = 1
		}
	}

	result := make([]rune, 0)

	for k, v := range counts {
		if v > 1 {
			result = append(result, k)
		}
	}

	return result
}