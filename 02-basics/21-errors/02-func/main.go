package main

import (
	"errors"
	"fmt"
)

func main() {
	new, err := ChangeSign(0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("The new number is %d\n", new)
	}
}

func ChangeSign(number int) (int, error) {
	var new int

	if number == 0 {
		return new, errors.New("the number cannot be 0")
	}

	new = number * -1

	return new, nil
}