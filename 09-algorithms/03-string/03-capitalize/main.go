package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "this is Gopher!"
	capitalized := Capitalize(str)

	fmt.Printf("%s -> %s\n", str, capitalized)
}

func Capitalize(str string) string {
	builder := strings.Builder{}

	builder.Grow(len(str))

	spaceFlag := true

	for _, c := range str {
		if spaceFlag {
			builder.WriteRune(UpperChar(c))
		} else {
			builder.WriteRune(c)
		}

		spaceFlag = c == ' '
	}

	return builder.String()
}

func UpperChar(c rune) rune {
	diff := 'a' - 'A'

	if 'a' <= c && c <= 'z' {
		return c - diff
	}

	return c
}