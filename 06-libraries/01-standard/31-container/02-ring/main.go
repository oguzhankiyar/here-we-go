package main

import (
	"container/ring"
	"fmt"
)

func main() {
	Sample("Ring", Ring)
}

func Ring() {
	values := ring.New(5)

	for i := 0; i < values.Len(); i++ {
		values.Value = i + 1
		values = values.Next()
	}

	values.Do(func(x interface{}) {
		fmt.Print(x)
	})

	fmt.Println()
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}