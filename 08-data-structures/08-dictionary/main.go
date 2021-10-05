package main

import "fmt"

func main() {
	dict := New()
	dict.Set("id", 1)
	dict.Set("name", "gopher")

	fmt.Println(dict.Keys())
	fmt.Println(dict.Values())
	fmt.Println()
	fmt.Println("id:", dict.Get("id"))
	fmt.Println("name:", dict.Get("name"))
}

type Dictionary struct {
	Items map[interface{}]interface{}
}

func New() *Dictionary {
	return &Dictionary{
		Items: make(map[interface{}]interface{}),
	}
}

func (dict *Dictionary) Get(k interface{}) interface{} {
	return dict.Items[k]
}

func (dict *Dictionary) Set(k interface{}, v interface{}) {
	if dict.Items == nil {
		dict.Items = make(map[interface{}]interface{})
	}
	dict.Items[k] = v
}

func (dict *Dictionary) Del(k interface{}) bool {
	_, ok := dict.Items[k]
	if ok {
		delete(dict.Items, k)
	}
	return ok
}

func (dict *Dictionary) Contains(k interface{}) bool {
	_, ok := dict.Items[k]
	return ok
}

func (dict *Dictionary) Clear() {
	dict.Items = make(map[interface{}]interface{})
}

func (dict *Dictionary) Size() int {
	return len(dict.Items)
}

func (dict *Dictionary) Keys() []interface{} {
	keys := make([]interface{}, 0)
	for i := range dict.Items {
		keys = append(keys, i)
	}
	return keys
}

func (dict *Dictionary) Values() []interface{} {
	values := make([]interface{}, 0)
	for i := range dict.Items {
		values = append(values, dict.Items[i])
	}
	return values
}