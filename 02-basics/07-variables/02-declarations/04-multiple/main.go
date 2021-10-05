package main

import "fmt"

func main() {
	var (
		firstName string = "Gopher"
		lastName string = "Go"
	)
	fmt.Println(firstName, lastName)

	var ageMonth, ageYear int = 5, 1993
	fmt.Println(ageMonth, ageYear)
}
