package main

import (
	"fmt"

	"external-sample/pkg/clients"
)

func main() {
	// The pkg folder is special
	// It can be used externally

	a := 3
	b := 5
	sum := clients.Sum(a, b)
	fmt.Printf("%d + %d = %d\n", a, b, sum)
}