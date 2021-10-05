package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i > 10 {
			break
		}

		if i % 2 == 0 {
			continue
		}

		fmt.Printf("%d ", i)
	}
}