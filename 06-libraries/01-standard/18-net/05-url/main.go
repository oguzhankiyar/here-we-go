package main

import (
	"fmt"
	"net/url"
)

func main() {
	Sample("Parse", Parse)
	Sample("ParseQuery", ParseQuery)
}

func Parse() {
	fn := func(str string) {
		result, err := url.Parse(str)
		if err != nil {
			fmt.Println("error:", err)
		}

		fmt.Println("-", str)
		fmt.Println("Scheme:", result.Scheme)
		fmt.Println("Host:", result.Host)
		fmt.Println("Port:", result.Port())
		fmt.Println("Path:", result.Path)
		fmt.Println("Query:", result.Query())
		fmt.Println("User:", result.User)

		fmt.Println()
	}

	fn("https://golang.org")
	fn("tcp://127.0.0.1:8080")
	fn("https://user1:pass1@golang.org")
	fn("https://golang.org/pkg/net/url")
	fn("https://example.com?search=golang")
}

func ParseQuery() {
	fn := func(str string) {
		result, err := url.ParseQuery(str)
		if err != nil {
			fmt.Println("error:", err)
		}

		fmt.Println("-", str)

		fmt.Println(result)

		fmt.Println()
	}

	fn("x=1&y=2") // map[x:[1] y:[2]]
	fn("x=&y=2") // map[x:[] y:[2]]
	fn("x=1&y=2&y=3;z") // map[x:[1] y:[2 3] z:[]]
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}