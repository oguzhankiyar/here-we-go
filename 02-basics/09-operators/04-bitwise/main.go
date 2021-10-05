package main

import "fmt"

func main() {
	a, b := 0, 1

	and := a & b
	fmt.Printf("a & b => %d\n", and)

	or := a | b
	fmt.Printf("a | b => %d\n", or)

	xor := a ^ b
	fmt.Printf("a ^ b => %d\n", xor)

	leftShift := a << b
	fmt.Printf("a << b => %d\n", leftShift)

	rightShift := a >> b
	fmt.Printf("a >> b => %d\n", rightShift)
}