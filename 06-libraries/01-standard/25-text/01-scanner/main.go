package main

import (
	"fmt"
	"strings"
	"text/scanner"
)

func main() {
	Sample("Next", Next)
	Sample("Scan", Scan)
}

func Next() {
	reader := strings.NewReader("Hello, there!")

	scn := scanner.Scanner{}

	scn.Init(reader)

	var items []string

	for {
		rn := scn.Next()

		if rn == scanner.EOF {
			break
		}

		str := scanner.TokenString(rn)

		items = append(items, str)
	}

	fmt.Println(strings.Join(items, ", "))
}

func Scan() {
	reader := strings.NewReader("Hello, there!")

	scn := scanner.Scanner{}

	scn.Init(reader)

	var items []string

	for {
		rn := scn.Scan()

		if rn == scanner.EOF {
			break
		}

		str := scanner.TokenString(rn)

		items = append(items, str)
	}

	fmt.Println(strings.Join(items, ", "))
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}