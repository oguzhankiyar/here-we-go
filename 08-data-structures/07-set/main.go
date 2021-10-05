package main

import "fmt"

func main() {
	set := New()
	set.Add(1)
	set.Add(2)
	set.Add(1)
	set.Add(3)
	set.Delete(2)
	set.Display()
}

type Set struct {
	Items map[interface{}]bool
}

func New() *Set {
	return &Set{make(map[interface{}]bool)}
}

func (set *Set) Add(data interface{}) *Set {
	if set.Items == nil {
		set.Items = make(map[interface{}]bool)
	}

	_, ok := set.Items[data]
	if !ok {
		set.Items[data] = true
	}

	return set
}

func (set *Set) Clear() {
	set.Items = make(map[interface{}]bool)
}

func (set *Set) Delete(data interface{}) bool {
	_, ok := set.Items[data]
	if ok {
		delete(set.Items, data)
	}
	return ok
}

func (set *Set) Has(data interface{}) bool {
	_, ok := set.Items[data]
	return ok
}

func (set *Set) Display() {
	for i := range set.Items {
		fmt.Println(i)
	}
}

func (set *Set) Size() int {
	return len(set.Items)
}