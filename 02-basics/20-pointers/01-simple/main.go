package main

import "fmt"

func main() {
	var num int = 5
	fmt.Printf("num address: %x\n", &num)

	var point *int

	point = &num

	fmt.Printf("point address: %x\n", point)
	fmt.Printf("point value: %d\n", *point)
}