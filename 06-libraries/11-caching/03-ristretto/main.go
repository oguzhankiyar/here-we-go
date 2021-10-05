package main

import (
	"fmt"

	"github.com/dgraph-io/ristretto"
)

func main() {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,
		MaxCost:     1 << 30,
		BufferItems: 64,
	})
	if err != nil {
		panic(err)
	}

	cache.Set("key", "my-value", 1)

	cache.Wait()

	value, found := cache.Get("key")
	if !found {
		panic("missing value")
	}

	fmt.Println(value)

	cache.Del("key")
}