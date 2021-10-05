package main

import (
	"fmt"
	"strings"
)

func main() {
	fn := func(str string) {
		reversed := Reverse(str)
		fmt.Printf("%s -> %s\n", str, reversed)
	}

	fn("gopher")
	fn("GoLang!")
}

func Reverse(str string) string {
	var builder strings.Builder

	builder.Grow(len(str))

	for i := len(str) - 1; i >= 0; i-- {
		c := str[i]
		builder.WriteByte(c)
	}

	return builder.String()
}