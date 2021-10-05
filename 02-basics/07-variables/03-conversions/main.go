package main

import "fmt"

func main() {
	var number int = 5
	var floated float32 = float32(number)

	fmt.Println("number:", number, "floated:", floated)

	var a int = 65 // byte value for A char
	var str string = string(a)
	fmt.Println("str", str)
}