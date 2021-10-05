package main

import (
	"fmt"
	"strings"
)

func main() {
	fn := func(str string) {
		upper := Lower(str)
		fmt.Printf("%s -> %s\n", str, upper)
	}

	fn("This")
	fn("is")
	fn("GOPHER!")
}

func Lower(str string) string {
	builder := strings.Builder{}

	builder.Grow(len(str))

	diff := 'a' - 'A'

	for _, c := range str {
		if 'A' <= c && c <= 'Z' {
			builder.WriteString(string(c + diff))
		} else {
			builder.WriteString(string(c))
		}
	}

	return builder.String()
}