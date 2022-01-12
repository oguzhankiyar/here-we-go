package main

import "fmt"

func main() {
	PrintFloatType(4.2)
}

type FloatType interface {
	float32 | float64
}

func PrintFloatType[T FloatType](v T) {
	switch any(v).(type) {
	case float32:
		fmt.Println("float32")
	case float64:
		fmt.Println("float64")
	default:
		fmt.Println("unexpected")
	}
}