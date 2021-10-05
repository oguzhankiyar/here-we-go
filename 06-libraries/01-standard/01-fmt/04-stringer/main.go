package main

import (
	"fmt"
)

type Item struct {
	Id 		int
	Name	string
}

func (i Item) String() string {
	return fmt.Sprintf("[#%d] %s", i.Id, i.Name)
}

func (i Item) GoString() string {
	return fmt.Sprintf("{ id:%d, name:%q }", i.Id, i.Name)
}

func main() {
	Sample("Stringer", Stringer)
	Sample("GoStringer", GoStringer)
}

func Stringer() {
	// fmt package has Stringer interface to write as string
	// to override fmt string process, you should write "func String() string" to your struct
	/*
		type Stringer interface {
			String() string
		}
	*/

	item := Item{1, "Item 1"}

	fmt.Printf("String(): %v\n", item)
}

func GoStringer() {
	// fmt package has GoStringer interface to write as string
	// to override fmt go string process, you should write "func GoString() string" to your struct
	/*
		type GoStringer interface {
			GoString() string
		}
	*/

	item := Item{2, "Item 2"}

	fmt.Printf("GoString(): %#v\n", item)
}

func Sample(name string, fn func())  {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}