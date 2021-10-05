package main

import (
	"fmt"
	"math"
)

func main() {
	perimeter, area := CalculateCircle(3)

	fmt.Printf("Perimeter: %v, Area: %v\n", perimeter, area)
}

func CalculateCircle(round float64) (float64, float64) {
	perimeter := math.Pi * round * 2
	area := math.Pi * math.Pow(round, 2)

	return perimeter, area
}