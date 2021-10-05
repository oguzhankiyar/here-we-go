package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	name := "Gopher!"
	length := len(name)
	fmt.Printf("len('%s') = %d\n", name, length)

	// Unicode
	name = "Oğuzhan"
	length = len(name)
	fmt.Printf("len('%s') = %d\n", name, length)

	// With utf8
	length = utf8.RuneCountInString(name)
	fmt.Printf("len('%s') = %d\n", name, length)
}
