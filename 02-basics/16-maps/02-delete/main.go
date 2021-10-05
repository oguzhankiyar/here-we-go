package main

import "fmt"

func main() {
	dict := map[string]string{
		"name": "Gopher",
		"age": "1",
		"temp": "tmp",
	}
	fmt.Println(dict)

	delete(dict, "temp")
	fmt.Println(dict)
}