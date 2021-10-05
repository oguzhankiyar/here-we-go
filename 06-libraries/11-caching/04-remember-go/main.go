package main

import (
	"fmt"
	"time"

	"github.com/rocketlaunchr/remember-go"
	"github.com/rocketlaunchr/remember-go/memory"
)

func main() {
	Sample("Simple", Simple)
	Sample("Complex", Complex)

}

func Simple() {
	ms := memory.NewMemoryStore(10 * time.Minute)

	err := ms.Set("key1", time.Hour, "value1")
	if err != nil {
		panic(err)
	}

	value, _, err := ms.Get("key1")
	if err != nil {
		panic(err)
	}
	fmt.Println("key1:", value)
}

func Complex() {
	ms := memory.NewMemoryStore(10 * time.Minute)

	type Key struct {
		Search string
		Page   int `json:"page"`
	}

	key1 := remember.CreateKeyStruct(Key{"golang", 2})
	key2 :=  remember.CreateKey(false, "-", "search-x-y", "search", "golang", 2)

	fmt.Println(key1)
	fmt.Println(key2)

	err := ms.Set(key1, time.Hour, "search-value")
	if err != nil {
		panic(err)
	}

	value, _, err := ms.Get(key1)
	if err != nil {
		panic(err)
	}
	fmt.Println(key1, ":", value)
}

func Sample(name string, fn func())  {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}