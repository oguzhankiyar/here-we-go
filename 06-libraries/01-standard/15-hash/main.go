package main

import (
	"fmt"
	"hash/fnv"
)

func main() {
	Sample("Fnv", Fnv)
}

func Fnv() {
	fn := func(str string) {
		h := fnv.New32a()
		h.Write([]byte(str))
		fmt.Printf("%s -> %v\n", str, h.Sum32())
	}

	fn("Hello, Gopher")
	fn("Hello, Gopher!")
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}
