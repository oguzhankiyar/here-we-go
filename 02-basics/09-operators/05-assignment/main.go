package main

import "fmt"

func main() {
	a, b := 10, 2

	var number int

	number = a
	fmt.Printf("number = a => %d\n", number)

	number += b
	fmt.Printf("number += b => %d\n", number)

	number += b
	fmt.Printf("number -= b => %d\n", number)

	number *= b
	fmt.Printf("number *= b => %d\n", number)

	number /= b
	fmt.Printf("number /= b => %d\n", number)

	number %= b
	fmt.Printf("number %%= b => %d\n", number)

	number <<= b
	fmt.Printf("number %%= b => %d\n", number)

	number >>= b
	fmt.Printf("number %%= b => %d\n", number)

	number &= b
	fmt.Printf("number &= b => %d\n", number)

	number |= b
	fmt.Printf("number |= b => %d\n", number)

	number ^= b
	fmt.Printf("number ^= b => %d\n", number)
}