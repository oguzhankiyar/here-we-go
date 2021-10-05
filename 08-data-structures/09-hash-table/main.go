package main

import (
	"fmt"
)

func main() {
	ht := New()
	ht.Set("id", 1)
	ht.Set("id", 2)
	fmt.Println(ht.Get("id"))
	ht.Del("id")
	fmt.Println(ht.Get("id"))
}

type HashTable struct {
	Items map[int]interface{}
}

func New() *HashTable {
	return &HashTable{
		Items: make(map[int]interface{}),
	}
}

func (ht *HashTable) Get(key interface{}) interface{} {
	i := Hash(key)
	return ht.Items[i]
}

func (ht *HashTable) Set(key interface{}, value interface{}) {
	i := Hash(key)
	if ht.Items == nil {
		ht.Items = make(map[int]interface{})
	}
	ht.Items[i] = value
}

func (ht *HashTable) Del(key interface{}) {
	i := Hash(key)
	delete(ht.Items, i)
}

func (ht *HashTable) Size() int {
	return len(ht.Items)
}

func Hash(k interface{}) int {
	key := fmt.Sprintf("%s", k)
	h := 0
	for i := 0; i < len(key); i++ {
		h = 31 * h + int(key[i])
	}
	return h
}