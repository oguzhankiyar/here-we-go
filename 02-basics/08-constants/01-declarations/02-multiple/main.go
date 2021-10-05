package main

func main() {
	// Type 1
	const a, b = 1, 2

	// Type 2
	const (
		c = 1
		d = 2
	)

	// Repeating
	const (
		e = 1	// int -> 1
		f		// int -> 1
	)
}
