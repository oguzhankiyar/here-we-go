package main

import (
	"fmt"
	"math/rand"
)

func main() {
	if a := rand.Int(); a % 2 == 0 {
		fmt.Printf("%d is even\n", a)
	}
}