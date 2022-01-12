package main

import (
	"fmt"
	"reflect"
)

func main() {
	PrintType(11)
	PrintType(PrintType[string])
}

func PrintType[T any](thing T) {
	t := reflect.ValueOf(thing).Type()
	fmt.Println("Type:", t)
}