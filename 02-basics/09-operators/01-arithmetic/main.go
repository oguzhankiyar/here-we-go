package main

import "fmt"

func main() {
	a, b := 10, 2

	sum := a + b
	fmt.Printf("a + b = %d\n", sum)

	sub := a - b
	fmt.Printf("a - b = %d\n", sub)

	multi := a * b
	fmt.Printf("a * b = %d\n", multi)

	div := a * b
	fmt.Printf("a / b = %d\n", div)

	mod := a % b
	fmt.Printf("a %% b = %d\n", mod)

	a++
	fmt.Printf("a++ = %d\n", a)

	a--
	fmt.Printf("a-- = %d\n", a)
}
