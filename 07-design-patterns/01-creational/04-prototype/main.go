package main

import "fmt"

type Node interface {
	GetName()	string
	Clone() Node
}

type Item struct {
	name string
}

func (f Item) GetName() string {
	return f.name
}

func (f Item) Clone() Node {
	return Item{f.name + "_clone"}
}

func main() {
	item1 := Item{"item"}
	item2 := item1.Clone()
	item3 := item2.Clone()

	fmt.Println(item1.GetName())
	fmt.Println(item2.GetName())
	fmt.Println(item3.GetName())
}