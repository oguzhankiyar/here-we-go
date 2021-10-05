package main

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

func main() {
	c := cache.New(5*time.Minute, 10*time.Minute)

	c.Set("foo", "bar", cache.DefaultExpiration)

	c.Set("baz", 42, cache.NoExpiration)

	foo, found := c.Get("foo")
	if found {
		fmt.Println(foo)
	}

	if x, found := c.Get("foo"); found {
		foo := x.(string)
		fmt.Println(foo)
	}

	type Data struct {
		Value string
	}

	data := Data{
		Value: "val-1",
	}

	c.Set("foo", &data, cache.DefaultExpiration)
	if x, found := c.Get("foo"); found {
		foo := x.(*Data)
		fmt.Println(foo.Value)
	}
}