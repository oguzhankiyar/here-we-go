package main

import (
	"container/list"
	"fmt"
)

func main() {
	Sample("Init", Init)
	Sample("Push", Push)
	Sample("Insert", Insert)
	Sample("Move", Move)
	Sample("Remove", Remove)
}

func Init() {
	values := list.New()

	values.PushBack("item 1")

	fmt.Println("before:", values.Len())

	values.Init()

	fmt.Println("after:", values.Len())
}

func Push() {
	values := list.New()

	values.PushBack("two")
	values.PushBack("three")
	values.PushFront("one")

	for item := values.Front(); item != nil; item = item.Next() {
		fmt.Println(item.Value)
	}
}

func Insert() {
	values := list.New()

	item1 := "item 1"

	fmt.Print("init: ")

	el := values.PushFront(item1)

	for item := values.Front(); item != nil; item = item.Next() {
		fmt.Printf("%v ", item.Value)
	}

	fmt.Println()
	fmt.Print("before: ")

	values.InsertBefore("one", el)

	for item := values.Front(); item != nil; item = item.Next() {
		fmt.Printf("%v ", item.Value)
	}

	fmt.Println()
	fmt.Print("after: ")

	values.InsertAfter("two", el)

	for item := values.Front(); item != nil; item = item.Next() {
		fmt.Printf("%v ", item.Value)
	}

	fmt.Println()
}

func Move() {
	values := list.New()

	fourEl := values.PushFront(4)
	threeEl := values.PushFront(3)
	twoEl := values.PushFront(2)
	oneEl := values.PushFront(1)

	fmt.Print("init: ")

	for item := values.Front(); item != nil; item = item.Next() {
		fmt.Printf("%v ", item.Value)
	}

	values.MoveToFront(fourEl)
	values.MoveToBack(oneEl)
	values.MoveAfter(twoEl, threeEl)

	fmt.Println()
	fmt.Print("move: ")

	for item := values.Front(); item != nil; item = item.Next() {
		fmt.Printf("%v ", item.Value)
	}

	fmt.Println()
}

func Remove() {
	values := list.New()

	oneEl := values.PushBack(1)
	twoEl := values.PushBack(2)

	fmt.Print("init: ")

	for item := values.Front(); item != nil; item = item.Next() {
		fmt.Printf("%v ", item.Value)
	}

	values.Remove(oneEl)

	_ = twoEl

	fmt.Println()
	fmt.Print("remove: ")

	for item := values.Front(); item != nil; item = item.Next() {
		fmt.Printf("%v ", item.Value)
	}

	fmt.Println()
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}