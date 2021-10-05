package main

import "fmt"

func main() {
	var dict map[string]string

	// True
	if dict == nil {
		fmt.Println("dict is nil")
	}

	// Runtime error
	// dict["name"] = "Gopher"

	fmt.Printf("dict['name'] = %#v\n", dict["name"])
	fmt.Printf("len(dict) = %d\n", len(dict))

	dict = map[string]string{
		"name": "Gopher",
	}

	dict["age"] = "1"

	fmt.Printf("dict['name'] = %#v\n", dict["name"])
	fmt.Printf("dict['age'] = %#v\n", dict["age"])
	fmt.Printf("len(dict) = %d\n", len(dict))
}