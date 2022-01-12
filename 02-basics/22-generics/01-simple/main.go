package main

import (
	"fmt"
)

func main() {
	Say("hey")
	Say(10)
	Say(false)
}

func Say[T any](thing T) {
	fmt.Println(thing)
}