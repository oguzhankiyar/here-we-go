package main

import "fmt"

func main() {
	xox := [4][4]int{
		{ 0, 1, 0, 0 },
		{ 0, 0, 1, 1 },
		{ 1, 1, 0, 0 },
		{ 0, 0, 1, 0 },
	}

	fmt.Println(xox)

	for i, row := range xox {
		for j, column := range row {
			fmt.Printf("(%d, %d) = %d\n", i, j, column)
		}
	}
}