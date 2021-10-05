package main

import "fmt"

func main() {
	// This semicolon is unnecessary
	var number int = 1;
	fmt.Println("Number is", number)

	// This semicolon is required to separate two declarations
	var hello string = "Hello"; var name string = "Gopher"
	fmt.Println(hello, name)
}
