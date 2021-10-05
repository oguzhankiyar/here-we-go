package main

import (
	"fmt"

	"github.com/VictoriaMetrics/fastcache"
)

func main() {
	fc := fastcache.New(1024)
	fc.Set([]byte("my-key"), []byte("my-value"))

	val := fc.Get(nil, []byte("my-key"))
	fmt.Printf("%s\n", val)
}