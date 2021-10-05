package main

import "fmt"

func main() {
	fn := func(str1 string, str2 string) {
		result := Anagram(str1, str2)
		fmt.Printf("%q - %q -> %v\n", str1, str2, result)
	}

	fn("silent", "listen")
	fn("heart", "earth")
	fn("go", "og")
	fn("go", "goo")
}

func Anagram(str1 string, str2 string) bool {
	str1Len := len(str1)
	str2Len := len(str2)

	if str1Len != str2Len {
		return false
	}

	str1Counts := Counts(str1)
	str2Counts := Counts(str2)

	for k, v := range str1Counts {
		if v != str2Counts[k] {
			return false
		}
	}

	return true
}

func Counts(str string) map[rune]int {
	counts := make(map[rune]int)

	diff := 'a' - 'A'

	for _, c := range str {
		if 'A' <= c && c <= 'Z' {
			counts[c + diff]++
		} else {
			counts[c]++
		}
	}

	return counts
}