package main

import "fmt"

func main() {
	fn := func(str string) {
		result := Pangram(str)
		fmt.Printf("%q -> %v\n", str, result)
	}

	fn("The quick brown fox jumps over the lazy dog.")
	fn("Gopher learns how to learn go.")
}

func Pangram(str string) bool {
	counts := make(map[rune]int)

	diff := 'a' - 'A'

	for _, c := range str {
		if 'A' <= c && c <= 'Z' {
			counts[c + diff]++
		} else {
			counts[c]++
		}
	}

	for c := 'a'; c < 'z'; c++ {
		if count, ok := counts[c]; !ok || count == 0 {
			return false
		}
	}

	return true
}