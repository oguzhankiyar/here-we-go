package main

import "fmt"

func main() {
	fn := func(str string) {
		result := Palindrome(str)
		fmt.Printf("%q -> %v\n", str, result)
	}

	fn("gopher")
	fn("sos")
	fn("gg")
	fn("i")
}

func Palindrome(str string) bool {
	for i := 0; i < len(str) / 2; i++ {
		if str[i] != str[len(str) - i - 1] {
			return false
		}
	}

	return true
}