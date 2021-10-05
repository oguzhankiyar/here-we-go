package main

import "fmt"

func main() {
	var dict map[string]string


	dict = map[string]string{
		"name": "Gopher",
		"lang": "en",
	}

	for key, value := range dict {
		fmt.Printf("%s: %s\n", key, value)
	}
}