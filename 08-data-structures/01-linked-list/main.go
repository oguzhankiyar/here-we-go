package main

import "fmt"

func main() {
	list := New()
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Delete(2)
	list.Insert(5)
	list.Insert(7)
	list.Delete(5)
	list.Insert(9)
	list.Display()
}

type LinkedList struct {
	Head *Node
}

type Node struct {
	Next *Node
	Data interface{}
}

func New() *LinkedList {
	return &LinkedList{nil}
}

func (list *LinkedList) IsEmpty() bool {
	return list.Head == nil
}

func (list *LinkedList) Find(data interface{}) int {
	index := 0
	current := list.Head

	for current != nil {
		if current.Data == data {
			return index
		}

		index++
		current = current.Next
	}

	return -1
}

func (list *LinkedList) Insert(data interface{}) {
	node := Node{
		Data: data,
	}

	if list.IsEmpty() {
		list.Head = &node
		return
	}

	current := list.Head

	for current.Next != nil {
		current = current.Next
	}

	current.Next = &node
}

func (list *LinkedList) Delete(data interface{}) {
	prev := list.Head
	current := list.Head

	for current != nil {
		if current.Data == data {
			prev.Next = current.Next
			return
		}

		prev = current
		current = current.Next
	}
}

func (list *LinkedList) Display() {
	current := list.Head

	for current != nil {
		fmt.Println(current.Data)

		current = current.Next
	}
}