package main

import (
	"fmt"

	"internal-sample/internal/utils"
	"internal-sample/internal/utils/calc"
)

func main() {
	// The internal folder is special
	// It cannot be used externally

	a := utils.Rand(0, 10)
	b := utils.Rand(0, 10)
	sum := calc.Sum(a, b)
	fmt.Printf("%d + %d = %d\n", a, b, sum)
}