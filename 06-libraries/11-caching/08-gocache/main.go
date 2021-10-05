package main

import (
	"fmt"
	"time"

	"github.com/allegro/bigcache"
	"github.com/eko/gocache/cache"
	"github.com/eko/gocache/store"
)

func main() {
	bgClient, _ := bigcache.NewBigCache(bigcache.DefaultConfig(5 * time.Minute))
	bgStore := store.NewBigcache(bgClient, nil)

	cacheManager := cache.New(bgStore)

	err := cacheManager.Set("my-key", []byte("my-value"), nil)
	if err != nil {
		panic(err)
	}

	value, err := cacheManager.Get("my-key")
	if err != nil {
		panic(err)
	}

	fmt.Printf("value: %s\n", value)
}