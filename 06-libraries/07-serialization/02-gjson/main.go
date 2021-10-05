package main

import (
	"fmt"
	"github.com/tidwall/gjson"
)

func main() {
	const json = `{"name":{"first":"Gopher","last":"Go"},"age":10}`

	value := gjson.Get(json, "name.first")

	fmt.Println(value.String())
}