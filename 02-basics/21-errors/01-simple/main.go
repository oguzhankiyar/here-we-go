package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("some error")
	fmt.Println(err)
}