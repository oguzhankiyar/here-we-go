package main

import (
	"fmt"
	"log"

	"github.com/teris-io/shortid"
)

func main() {
	sid, err := shortid.New(1, shortid.DefaultABC, 2342)
	if err != nil {
		log.Fatal(err)
	}

	var result string

	result, err = sid.Generate()
	fmt.Println("id:", result)
	result, err = sid.Generate()
	fmt.Println("id:", result)

	shortid.SetDefault(sid)

	result, err = shortid.Generate()
	fmt.Println("id:", result)
	result, err = shortid.Generate()
	fmt.Println("id:", result)
}