package main

import "fmt"

func main() {
	max := 5
	current := 1

	step:
		if current <= max {
			fmt.Printf("%d ", current)
			current++
			goto step
		}
}