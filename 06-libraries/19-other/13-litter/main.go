package main

import (
	"github.com/sanity-io/litter"
)

func main() {
	type Person struct {
		Name   string
		Age    int
		Parent *Person
	}

	litter.Dump(Person{
		Name:   "Bob",
		Age:    20,
		Parent: &Person{
			Name: "Jane",
			Age:  50,
		},
	})
}
