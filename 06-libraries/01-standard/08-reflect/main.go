package main

import (
	"fmt"
	"reflect"
)

func main() {
	Sample("Type", Type)
	Sample("TypeOf", TypeOf)
}

func Type() {

}

func TypeOf() {
	fn := func(value interface{}) {
		fmt.Printf("%#v -> %s\n", value, reflect.TypeOf(value))
	}

	type Post struct {
		Title string
	}

	fn("Hello")// string
	fn(100) // int
	fn(7.2) // float64
	fn(true) // bool
	fn(struct{}{}) // struct {}
	fn(Post{ "Post 1" }) // main.go.Post
}

func Sample(name string, fn func())  {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}