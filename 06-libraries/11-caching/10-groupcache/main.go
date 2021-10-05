package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/golang/groupcache"
)

func main() {
	store := map[string][]byte{
		"red":   []byte("#FF0000"),
		"green": []byte("#00FF00"),
		"blue":  []byte("#0000FF"),
	}

	group := groupcache.NewGroup("foobar", 64<<20, groupcache.GetterFunc(
		func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
			log.Println("looking up", key)
			v, ok := store[key]
			if !ok {
				return errors.New("color not found")
			}
			dest.SetBytes(v)
			return nil
		},
	))

	var b []byte
	err := group.Get(nil, "red", groupcache.AllocatingByteSliceSink(&b))
	if err != nil {
		return
	}
	fmt.Printf("%s\n", b)
}