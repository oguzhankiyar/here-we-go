package main

import "fmt"

func main() {
	list := NewList[string]("hello")
	mapped := list.Map(func(i string)any {
		return i + "s"
	})
	fmt.Println(mapped)
}

type List[T any] struct {
	items []T
}

func NewList[T any](items ...T) *List[T] {
	list := make([]T, 0)
	list = append(list, items...)

	return &List[T]{
		items: list,
	}
}

func (l *List[T]) Map(fn func(T)any) []any {
	result := make([]any, 0)
	for _, v := range l.items {
		result = append(result, fn(v))
	}
	return result
}