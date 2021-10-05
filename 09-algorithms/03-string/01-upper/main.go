package main

import (
	"fmt"
	"strings"
)

func main() {
	fn := func(str string) {
		upper := Upper(str)
		fmt.Printf("%s -> %s\n", str, upper)
	}

	fn("This")
	fn("is")
	fn("GOPHER!")
}

func Upper(str string) string {
	builder := strings.Builder{}

	builder.Grow(len(str))

	diff := 'a' - 'A'

	for _, c := range str {
		if 'a' <= c && c <= 'z' {
			builder.WriteString(string(c - diff))
		} else {
			builder.WriteString(string(c))
		}
	}

	return builder.String()
}