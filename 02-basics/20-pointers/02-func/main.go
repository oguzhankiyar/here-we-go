package main

import "fmt"

func main() {
	Sample("WithoutPointer", WithoutPointer)
	Sample("WithPointer", WithPointer)
}

type Item struct {
	Value string
}

func WithoutPointer() {
	item := Item{Value: "item-1"}

	fmt.Printf("BEFORE ChangeValueWithoutPointer: %s\n", item.Value)
	ChangeValueWithoutPointer(item)
	fmt.Printf("AFTER ChangeValueWithoutPointer: %s\n", item.Value)
}

func WithPointer() {
	item := Item{Value: "item-1"}

	fmt.Printf("BEFORE ChangeValueWithPointer: %s\n", item.Value)
	ChangeValueWithPointer(&item)
	fmt.Printf("AFTER ChangeValueWithPointer: %s\n", item.Value)
}

func ChangeValueWithoutPointer(item Item) {
	item.Value = "item-2"

	fmt.Printf("INSIDE ChangeValueWithoutPointer %s\n", item.Value)
}

func ChangeValueWithPointer(item *Item) {
	item.Value = "item-2"

	fmt.Printf("INSIDE ChangeValueWithPointer %s\n", item.Value)
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}