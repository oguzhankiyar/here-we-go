package main

import (
	"fmt"
	"os"

	"github.com/peterbourgon/diskv"
)

func main() {
	flatTransform := func(s string) []string { return []string{} }

	d := diskv.New(diskv.Options{
		BasePath:     os.TempDir(),
		Transform:    flatTransform,
		CacheSizeMax: 1024 * 1024,
	})

	key := "alpha"
	err := d.Write(key, []byte("hey"))
	if err != nil {
		panic(err)
	}

	value, _ := d.Read(key)
	fmt.Printf("%s\n", value)

	err = d.Erase(key)
	if err != nil {
		panic(err)
	}
}