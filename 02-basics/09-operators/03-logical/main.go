package main

import "fmt"

func main() {
	a, b := true, false

	and := a && b
	fmt.Printf("a && b => %v\n", and)

	or := a || b
	fmt.Printf("a || b => %v\n", or)

	not := !a
	fmt.Printf("!a => %v\n", not)
}