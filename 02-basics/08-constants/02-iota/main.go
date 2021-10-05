package main

import "fmt"

func main() {
	// Without iota
	const (
		jan = 1
		feb = 2
		mar = 3
		apr = 4
		may = 5
		jun = 6
		jul = 7
		aug = 8
		sep = 9
		oct = 10
		nov = 11
		dec = 12
	)
	fmt.Println(jan, feb, mar, apr, may, jun, jul, aug, sep, oct, nov, dec)

	// With iota
	const (
		mon = iota
		tue
		wed
		thu
		fri
		sat
		sun
	)
	fmt.Println(mon, tue, wed, thu, fri, sat, sun)

	// Expression
	const (
		created = iota + 1
		updated
		deleted
	)
	fmt.Println(created, updated, deleted)

	// Blank
	const (
		first = iota + 2	// 2
		_
		third				// 4
	)
	fmt.Println(first, third)
}