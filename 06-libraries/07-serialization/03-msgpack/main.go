package main

import (
	"fmt"
	"github.com/vmihailenco/msgpack/v5"
)

func main() {
	type Thing struct {
		Name string
	}

	b, err := msgpack.Marshal(&Thing{Name: "Gopher"})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	var thing Thing
	err = msgpack.Unmarshal(b, &thing)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println(thing.Name)
}