package main

import "fmt"

func main() {
	print := func (str string) {
		fmt.Println(str)
	}

	print("Go!")
}