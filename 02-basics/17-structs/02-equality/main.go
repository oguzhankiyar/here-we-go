package main

import "fmt"

func main() {
	type Bike struct {
		Brand	string
		Model 	string
	}

	scott := Bike{
		"Scott",
		"Speedster 20 Disc",
	}

	other := scott

	fmt.Printf("scott: %+v\n", scott)
	fmt.Printf("other: %+v\n", other)
	fmt.Printf("equal: %v\n", scott == other)

	fmt.Println()

	new := Bike{
		"Scott",
		"Speedster 20 Disc",
	}

	fmt.Printf("scott: %+v\n", scott)
	fmt.Printf("new: %+v\n", new)
	fmt.Printf("equal: %v\n", scott == new)
}