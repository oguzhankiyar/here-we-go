package main

func main() {
	// Type 1
	const num1 int = 5

	// Type 2
	const num2 = 6.2

	// You cannot change the value of a constant
	// num2 = 6.3

	// Divide by zero check at compile-time
	const a = 1
	const b = 0
	// const c = a / b

	// Cannot init with a non-const variable
	var d = 1
	// const e = d
	_ = d

	// Cannot init with runtime function
	// const f = rand.Int()
}