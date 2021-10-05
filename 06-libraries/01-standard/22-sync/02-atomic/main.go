package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	num := 5

	val := atomic.Value{}
	val.Store(num)

	load := val.Load()
	number := load.(int)
	fmt.Println(num, number)

	val.Store(10)

	load = val.Load()
	number = load.(int)
	fmt.Println(num, number)
}