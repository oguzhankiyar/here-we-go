package main

import (
	"fmt"

	"github.com/mitchellh/hashstructure"
)

func main() {
	type ComplexStruct struct {
		Name     string
		Age      uint
		Metadata map[string]interface{}
	}

	v := ComplexStruct{
		Name: "Gopher",
		Age:  10,
		Metadata: map[string]interface{}{
			"car":      true,
			"location": "California",
			"siblings": []string{"Bob", "John"},
		},
	}

	hash, err := hashstructure.Hash(v, hashstructure.FormatV2, nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d", hash)
}
