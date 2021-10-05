package main

import "math"

func main() {
	// With type
	const a int = 1

	// Cannot assign to float
	// const b float64 = a * math.Pi

	// Without type
	const c = 1

	// Can assign to float
	const d = c * math.Pi
}