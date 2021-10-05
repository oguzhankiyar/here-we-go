package main

import "fmt"

type Iterator interface {
	Current()	interface{}
	MoveNext() 	bool
}

type List struct {
	Items []string
}

func (l List) GetIterator() Iterator {
	return &ListIterator{l, -1}
}

type ListIterator struct {
	List	List
	Index	int
}

func (i ListIterator) Current() interface{} {
	return i.List.Items[i.Index]
}

func (i *ListIterator) MoveNext() bool {
	if i.Index + 1 >= len(i.List.Items) {
		return false
	}

	i.Index++

	return true
}

func main() {
	list := List{[]string{"one", "two", "three"}}
	iterator := list.GetIterator()

	for iterator.MoveNext() {
		fmt.Println(iterator.Current())
	}
}