package main

import "fmt"

func main() {
	fmt.Printf("%d + %d = %d\n",1, 2, sum(1, 2))
}

func sum(numbers ...int) (result int) {
	for _, v := range numbers {
		result += v
	}

	return
}